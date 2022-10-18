// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// CreateSnapshot implements createSnapshot operation.
	//
	// Creates a snapshot of the microVM state. The microVM should be in the `Paused` state.
	//
	// PUT /snapshot/create
	CreateSnapshot(ctx context.Context, req SnapshotCreateParams) (CreateSnapshotRes, error)
	// CreateSyncAction implements createSyncAction operation.
	//
	// Creates a synchronous action.
	//
	// PUT /actions
	CreateSyncAction(ctx context.Context, req InstanceActionInfo) (CreateSyncActionRes, error)
	// DescribeBalloonConfig implements describeBalloonConfig operation.
	//
	// Returns the current balloon device configuration.
	//
	// GET /balloon
	DescribeBalloonConfig(ctx context.Context) (DescribeBalloonConfigRes, error)
	// DescribeBalloonStats implements describeBalloonStats operation.
	//
	// Returns the latest balloon device statistics, only if enabled pre-boot.
	//
	// GET /balloon/statistics
	DescribeBalloonStats(ctx context.Context) (DescribeBalloonStatsRes, error)
	// DescribeInstance implements describeInstance operation.
	//
	// Returns general information about an instance.
	//
	// GET /
	DescribeInstance(ctx context.Context) (DescribeInstanceRes, error)
	// GetExportVmConfig implements getExportVmConfig operation.
	//
	// Gets configuration for all VM resources.
	//
	// GET /vm/config
	GetExportVmConfig(ctx context.Context) (GetExportVmConfigRes, error)
	// GetMachineConfiguration implements getMachineConfiguration operation.
	//
	// Gets the machine configuration of the VM. When called before the PUT operation, it will return the
	// default values for the vCPU count (=1), memory size (=128 MiB). By default Hyperthreading is
	// disabled and there is no CPU Template.
	//
	// GET /machine-config
	GetMachineConfiguration(ctx context.Context) (GetMachineConfigurationRes, error)
	// LoadSnapshot implements loadSnapshot operation.
	//
	// Loads the microVM state from a snapshot. Only accepted on a fresh Firecracker process (before
	// configuring any resource other than the Logger and Metrics).
	//
	// PUT /snapshot/load
	LoadSnapshot(ctx context.Context, req SnapshotLoadParams) (LoadSnapshotRes, error)
	// MmdsConfigPut implements PUT /mmds/config operation.
	//
	// Creates MMDS configuration to be used by the MMDS network stack.
	//
	// PUT /mmds/config
	MmdsConfigPut(ctx context.Context, req MmdsConfig) (MmdsConfigPutRes, error)
	// MmdsGet implements GET /mmds operation.
	//
	// Get the MMDS data store.
	//
	// GET /mmds
	MmdsGet(ctx context.Context) (MmdsGetRes, error)
	// MmdsPatch implements PATCH /mmds operation.
	//
	// Updates the MMDS data store.
	//
	// PATCH /mmds
	MmdsPatch(ctx context.Context, req *MmdsPatchReq) (MmdsPatchRes, error)
	// MmdsPut implements PUT /mmds operation.
	//
	// Creates a MMDS (Microvm Metadata Service) data store.
	//
	// PUT /mmds
	MmdsPut(ctx context.Context, req *MmdsPutReq) (MmdsPutRes, error)
	// PatchBalloon implements patchBalloon operation.
	//
	// Updates an existing balloon device, before or after machine startup. Will fail if update is not
	// possible.
	//
	// PATCH /balloon
	PatchBalloon(ctx context.Context, req BalloonUpdate) (PatchBalloonRes, error)
	// PatchBalloonStatsInterval implements patchBalloonStatsInterval operation.
	//
	// Updates an existing balloon device statistics interval, before or after machine startup. Will fail
	// if update is not possible.
	//
	// PATCH /balloon/statistics
	PatchBalloonStatsInterval(ctx context.Context, req BalloonStatsUpdate) (PatchBalloonStatsIntervalRes, error)
	// PatchGuestDriveByID implements patchGuestDriveByID operation.
	//
	// Updates the properties of the drive with the ID specified by drive_id path parameter. Will fail if
	// update is not possible.
	//
	// PATCH /drives/{drive_id}
	PatchGuestDriveByID(ctx context.Context, req PartialDrive, params PatchGuestDriveByIDParams) (PatchGuestDriveByIDRes, error)
	// PatchGuestNetworkInterfaceByID implements patchGuestNetworkInterfaceByID operation.
	//
	// Updates the rate limiters applied to a network interface.
	//
	// PATCH /network-interfaces/{iface_id}
	PatchGuestNetworkInterfaceByID(ctx context.Context, req PartialNetworkInterface, params PatchGuestNetworkInterfaceByIDParams) (PatchGuestNetworkInterfaceByIDRes, error)
	// PatchMachineConfiguration implements patchMachineConfiguration operation.
	//
	// Partially updates the Virtual Machine Configuration with the specified input. If any of the
	// parameters has an incorrect value, the whole update fails.
	//
	// PATCH /machine-config
	PatchMachineConfiguration(ctx context.Context, req OptMachineConfiguration) (PatchMachineConfigurationRes, error)
	// PatchVm implements patchVm operation.
	//
	// Sets the desired state (Paused or Resumed) for the microVM.
	//
	// PATCH /vm
	PatchVm(ctx context.Context, req VM) (PatchVmRes, error)
	// PutBalloon implements putBalloon operation.
	//
	// Creates a new balloon device if one does not already exist, otherwise updates it, before machine
	// startup. This will fail after machine startup. Will fail if update is not possible.
	//
	// PUT /balloon
	PutBalloon(ctx context.Context, req Balloon) (PutBalloonRes, error)
	// PutGuestBootSource implements putGuestBootSource operation.
	//
	// Creates new boot source if one does not already exist, otherwise updates it. Will fail if update
	// is not possible.
	//
	// PUT /boot-source
	PutGuestBootSource(ctx context.Context, req BootSource) (PutGuestBootSourceRes, error)
	// PutGuestDriveByID implements putGuestDriveByID operation.
	//
	// Creates new drive with ID specified by drive_id path parameter. If a drive with the specified ID
	// already exists, updates its state based on new input. Will fail if update is not possible.
	//
	// PUT /drives/{drive_id}
	PutGuestDriveByID(ctx context.Context, req Drive, params PutGuestDriveByIDParams) (PutGuestDriveByIDRes, error)
	// PutGuestNetworkInterfaceByID implements putGuestNetworkInterfaceByID operation.
	//
	// Creates new network interface with ID specified by iface_id path parameter.
	//
	// PUT /network-interfaces/{iface_id}
	PutGuestNetworkInterfaceByID(ctx context.Context, req NetworkInterface, params PutGuestNetworkInterfaceByIDParams) (PutGuestNetworkInterfaceByIDRes, error)
	// PutGuestVsock implements putGuestVsock operation.
	//
	// The first call creates the device with the configuration specified in body. Subsequent calls will
	// update the device configuration. May fail if update is not possible.
	//
	// PUT /vsock
	PutGuestVsock(ctx context.Context, req Vsock) (PutGuestVsockRes, error)
	// PutLogger implements putLogger operation.
	//
	// Initializes the logger by specifying a named pipe or a file for the logs output.
	//
	// PUT /logger
	PutLogger(ctx context.Context, req Logger) (PutLoggerRes, error)
	// PutMachineConfiguration implements putMachineConfiguration operation.
	//
	// Updates the Virtual Machine Configuration with the specified input. Firecracker starts with
	// default values for vCPU count (=1) and memory size (=128 MiB). With Hyperthreading enabled, the
	// vCPU count is restricted to be 1 or an even number, otherwise there are no restrictions regarding
	// the vCPU count. If any of the parameters has an incorrect value, the whole update fails.
	//
	// PUT /machine-config
	PutMachineConfiguration(ctx context.Context, req OptMachineConfiguration) (PutMachineConfigurationRes, error)
	// PutMetrics implements putMetrics operation.
	//
	// Initializes the metrics system by specifying a named pipe or a file for the metrics output.
	//
	// PUT /metrics
	PutMetrics(ctx context.Context, req Metrics) (PutMetricsRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...Option) (*Server, error) {
	s, err := newConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}
