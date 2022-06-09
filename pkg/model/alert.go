package model

import (
	"time"
)

type Alert struct {
	UUID          string     `json:"uuid" gorm:"primarykey"`
	From          string     `json:"from" get:"true"` //
	HostUUID      string     `json:"host_uuid"`
	Host          string     `json:"host" get:"true"`
	AlertType     string     `json:"alert_type" get:"true"`
	Count         uint       `json:"count" get:"true"`
	Date          time.Time  `json:"date" get:"true"`
	FromAutomater bool       `json:"from_automater"`
	IsPipeline    bool       `json:"is_pipeline"`
	PipelineUUID  string     `json:"pipeline_uuid"`
	JobUUID       string     `json:"job_uuid"`
	Messages      []*Message `json:"messages" gorm:"constraint:OnDelete:CASCADE"`
}

type Message struct {
	UUID      string    `json:"uuid" gorm:"primarykey"  get:"true" post:"true" patch:"true" delete:"true"`
	Title     string    `json:"title,omitempty" get:"true"`
	Message   string    `json:"message,omitempty" get:"true"`
	Type      string    `json:"type,omitempty" get:"true"`
	Date      time.Time `json:"date,omitempty" get:"true"`
	AlertUUID string    `json:"alert_uuid,omitempty" get:"true" post:"true" patch:"true" gorm:"TYPE:string REFERENCES alerts;"`
}
