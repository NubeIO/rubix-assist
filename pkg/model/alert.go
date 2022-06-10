package model

import (
	"time"
)

type Task struct {
	UUID          string         `json:"uuid" gorm:"primarykey"`
	From          string         `json:"from"`
	HostUUID      string         `json:"host_uuid"`
	Host          string         `json:"host,omitempty"`
	TaskType      string         `json:"task_type"` //edgeapi.TaskType
	Count         uint           `json:"count"`
	Date          time.Time      `json:"date"`
	FromAutomater bool           `json:"from_automater"`
	IsPipeline    bool           `json:"is_pipeline,omitempty"`
	PipelineUUID  string         `json:"pipeline_uuid,omitempty"`
	IsJob         bool           `json:"is_job,omitempty"`
	JobUUID       string         `json:"job_uuid,omitempty"`
	Transactions  []*Transaction `json:"tasks" gorm:"constraint:OnDelete:CASCADE"`
}

type Transaction struct {
	UUID        string    `json:"uuid" gorm:"primarykey"`
	Title       string    `json:"title,omitempty"`
	Transaction string    `json:"message,omitempty"`
	Type        string    `json:"type,omitempty"`
	Date        time.Time `json:"date,omitempty"`
	TaskUUID    string    `json:"task_uuid,omitempty" gorm:"TYPE:string REFERENCES tasks;"`
}
