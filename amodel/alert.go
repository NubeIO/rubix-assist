package amodel

import (
	"time"
)

type Alert struct {
	UUID       string     `json:"uuid" gorm:"primarykey"`
	ParentUUID string     `json:"parent_uuid"`
	EntityType string     `json:"entity_type"`
	Type       string     `json:"type"`
	Status     string     `json:"status"`
	Message    string     `json:"message,omitempty"`
	Notes      string     `json:"notes,omitempty"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
}

type AlertClosed struct {
	Alert
	ClosedAt *time.Time `json:"closed_at,omitempty"`
}
