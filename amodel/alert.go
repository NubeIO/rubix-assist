package amodel

import (
	"time"
)

type Alert struct {
	UUID        string     `json:"uuid" gorm:"primarykey"`
	HostUUID    string     `json:"host_uuid"`
	EntityType  string     `json:"entity_type"`           // Device
	EntityUUID  string     `json:"entity_uuid,omitempty"` // dev_abc123
	Type        string     `json:"type"`                  // Ping
	Status      string     `json:"status"`                // Active
	Severity    string     `json:"severity"`              // Crucial
	Message     string     `json:"message,omitempty"`     // ping failed
	Notes       string     `json:"notes,omitempty"`       // notes by the user
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	LastUpdated *time.Time `json:"last_updated,omitempty"`
}

type AlertClosed struct {
	Alert
	ClosedAt *time.Time `json:"closed_at,omitempty"`
}
