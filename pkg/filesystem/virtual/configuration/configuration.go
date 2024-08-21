package configuration

import (
	"github.com/buildbarn/bb-remote-execution/pkg/filesystem/virtual"
	"github.com/buildbarn/bb-remote-execution/pkg/filesystem/virtual/nfsv4"
	pb "github.com/buildbarn/bb-remote-execution/pkg/proto/configuration/filesystem/virtual"
	"github.com/buildbarn/bb-storage/pkg/clock"
	"github.com/buildbarn/bb-storage/pkg/eviction"
	"github.com/buildbarn/bb-storage/pkg/program"
	"github.com/buildbarn/bb-storage/pkg/random"
	"github.com/buildbarn/bb-storage/pkg/util"
	nfsv4_xdr "github.com/buildbarn/go-xdr/pkg/protocols/nfsv4"
	"github.com/buildbarn/go-xdr/pkg/rpcserver"
	"github.com/jmespath/go-jmespath"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Mount of a virtual file system that has been created using
// NewMountFromConfiguration(), but that hasn't been exposed to the
// kernel or network yet. Before calling Expose(), the caller has the
// possibility to construct a root directory.
type Mount interface {
	Expose(terminationGroup program.Group, rootDirectory virtual.Directory) error
}

type fuseMount struct {
	mountPath       string
	configuration   *pb.FUSEMountConfiguration
	handleAllocator *virtual.FUSEStatefulHandleAllocator
	fsName          string
}

type nfsv4Mount struct {
	mountPath                        string
	configuration                    *pb.NFSv4MountConfiguration
	handleAllocator                  *virtual.NFSStatefulHandleAllocator
	authenticator                    rpcserver.Authenticator
	fsName                           string
	rootDirectoryAttributeCaching    AttributeCachingDuration
	childDirectoriesAttributeCaching AttributeCachingDuration
	leavesAttributeCaching           AttributeCachingDuration
}

func (m *nfsv4Mount) Expose(terminationGroup program.Group, rootDirectory virtual.Directory) error {
	// Random values that the client can use to identify the server, or
	// detect that the server has been restarted and lost all state.
	var soMajorID [8]byte
	random.FastThreadSafeGenerator.Read(soMajorID[:])
	serverOwner := nfsv4_xdr.ServerOwner4{
		SoMinorId: random.FastThreadSafeGenerator.Uint64(),
		SoMajorId: soMajorID[:],
	}
	var serverScope [8]byte
	random.FastThreadSafeGenerator.Read(serverScope[:])
	var rebootVerifier nfsv4_xdr.Verifier4
	random.FastThreadSafeGenerator.Read(rebootVerifier[:])
	var stateIDOtherPrefix [4]byte
	random.FastThreadSafeGenerator.Read(stateIDOtherPrefix[:])

	enforcedLeaseTime := m.configuration.EnforcedLeaseTime
	if err := enforcedLeaseTime.CheckValid(); err != nil {
		return util.StatusWrap(err, "Invalid enforced lease time")
	}
	announcedLeaseTime := m.configuration.AnnouncedLeaseTime
	if err := announcedLeaseTime.CheckValid(); err != nil {
		return util.StatusWrap(err, "Invalid announced lease time")
	}

	// Create a single server that is capable of accepting both
	// NFSv4.0 and NFSv4.1 requests.
	openedFilesPool := nfsv4.NewOpenedFilesPool(m.handleAllocator.ResolveHandle)
	program := nfsv4.NewMinorVersionFallbackProgram([]nfsv4_xdr.Nfs4Program{
		nfsv4.NewNFS41Program(
			rootDirectory,
			openedFilesPool,
			serverOwner,
			serverScope[:],
			// TODO: Should any of these parameters be configurable?
			&nfsv4_xdr.ChannelAttrs4{
				CaMaxrequestsize:        2 * 1024 * 1024,
				CaMaxresponsesize:       2 * 1024 * 1024,
				CaMaxresponsesizeCached: 64 * 1024,
				CaMaxoperations:         1000,
				CaMaxrequests:           100,
			},
			random.NewFastSingleThreadedGenerator(),
			rebootVerifier,
			clock.SystemClock,
			enforcedLeaseTime.AsDuration(),
			announcedLeaseTime.AsDuration(),
		),
		nfsv4.NewNFS40Program(
			rootDirectory,
			openedFilesPool,
			random.NewFastSingleThreadedGenerator(),
			rebootVerifier,
			stateIDOtherPrefix,
			clock.SystemClock,
			enforcedLeaseTime.AsDuration(),
			announcedLeaseTime.AsDuration(),
		),
	})

	// Create an RPC server that offers the NFSv4 program.
	rpcServer := rpcserver.NewServer(map[uint32]rpcserver.Service{
		nfsv4_xdr.NFS4_PROGRAM_PROGRAM_NUMBER: nfsv4_xdr.NewNfs4ProgramService(
			nfsv4.NewMetricsProgram(program),
		),
	}, m.authenticator)

	return m.mount(terminationGroup, rpcServer)
}

// NewMountFromConfiguration creates a new FUSE mount based on options
// specified in a configuration message and starts processing of
// incoming requests.
func NewMountFromConfiguration(configuration *pb.MountConfiguration, fsName string, rootDirectoryAttributeCaching, childDirectoriesAttributeCaching, leavesAttributeCaching AttributeCachingDuration) (Mount, virtual.StatefulHandleAllocator, error) {
	if configuration == nil {
		return nil, nil, status.Error(codes.InvalidArgument, "No mount configuration provided")
	}

	switch backend := configuration.Backend.(type) {
	case *pb.MountConfiguration_Fuse:
		handleAllocator := virtual.NewFUSEHandleAllocator(random.FastThreadSafeGenerator)
		return &fuseMount{
			mountPath:       configuration.MountPath,
			configuration:   backend.Fuse,
			handleAllocator: handleAllocator,
			fsName:          fsName,
		}, handleAllocator, nil
	case *pb.MountConfiguration_Nfsv4:
		handleAllocator := virtual.NewNFSHandleAllocator(random.NewFastSingleThreadedGenerator())

		authenticator := rpcserver.AllowAuthenticator
		if systemAuthentication := backend.Nfsv4.SystemAuthentication; systemAuthentication != nil {
			compiledExpression, err := jmespath.Compile(systemAuthentication.MetadataJmespathExpression)
			if err != nil {
				return nil, nil, util.StatusWrap(err, "Failed to compile system authentication metadata JMESPath expression")
			}
			evictionSet, err := eviction.NewSetFromConfiguration[nfsv4.SystemAuthenticatorCacheKey](systemAuthentication.CacheReplacementPolicy)
			if err != nil {
				return nil, nil, util.StatusWrap(err, "Failed to create system authentication eviction set")
			}
			authenticator = nfsv4.NewSystemAuthenticator(
				compiledExpression,
				int(systemAuthentication.MaximumCacheSize),
				eviction.NewMetricsSet(evictionSet, "SystemAuthenticator"))
		}

		return &nfsv4Mount{
			mountPath:                        configuration.MountPath,
			configuration:                    backend.Nfsv4,
			handleAllocator:                  handleAllocator,
			authenticator:                    authenticator,
			fsName:                           fsName,
			rootDirectoryAttributeCaching:    rootDirectoryAttributeCaching,
			childDirectoriesAttributeCaching: childDirectoriesAttributeCaching,
			leavesAttributeCaching:           leavesAttributeCaching,
		}, handleAllocator, nil
	default:
		return nil, nil, status.Error(codes.InvalidArgument, "No virtual file system backend configuration provided")
	}
}
