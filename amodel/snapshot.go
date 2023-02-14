package amodel

type CreateStatus string

const (
	Creating CreateStatus = "Creating"
)

type RestoreStatus string

const (
	Restoring RestoreStatus = "Restoring"
)

type CreateSnapshotStatus struct {
	UUID     string       `json:"-"`
	HostUUID string       `json:"host_uuid"`
	Status   CreateStatus `json:"status"`
}

type RestoreSnapshotStatus struct {
	UUID     string        `json:"-"`
	HostUUID string        `json:"host_uuid"`
	Status   RestoreStatus `json:"status"`
}

type SnapshotStatus struct {
	CreateStatus  []*CreateSnapshotStatus  `json:"create_status"`
	RestoreStatus []*RestoreSnapshotStatus `json:"restore_status"`
}
