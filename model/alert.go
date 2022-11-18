package model

import (
	"gorm.io/datatypes"
	"time"
)

type Task struct {
	UUID          string         `json:"uuid" gorm:"primarykey"`
	Status        string         `json:"status"`
	UserUUID      string         `json:"user_uuid"`
	HostUUID      string         `json:"host_uuid"`
	HostName      string         `json:"host_name,omitempty"`
	Type          string         `json:"type"`
	Count         uint           `json:"count"`
	CreatedAt     time.Time      `json:"date"`
	FromAutomater bool           `json:"from_automater"`
	IsPipeline    bool           `json:"is_pipeline,omitempty"`
	PipelineUUID  string         `json:"pipeline_uuid,omitempty"`
	IsJob         bool           `json:"is_job,omitempty"`
	JobUUID       string         `json:"job_uuid,omitempty"`
	Message       string         `json:"message,omitempty"`
	Data          datatypes.JSON `json:"data,omitempty"`
	Transactions  []*Transaction `json:"tasks" gorm:"constraint:OnDelete:CASCADE"`
}

type Transaction struct {
	UUID        string         `json:"uuid" gorm:"primarykey"`
	Status      string         `json:"status"`
	TaskType    string         `json:"type,omitempty"`
	SubTaskType string         `json:"sub_task_type"`
	Message     string         `json:"message,omitempty"`
	Data        datatypes.JSON `json:"data,omitempty"`
	TaskUUID    string         `json:"task_uuid,omitempty" gorm:"TYPE:string REFERENCES tasks;"`

	// JobID is the auto-generated pipeline identifier in UUID4 format.
	JobID string `json:"job_id"`

	IsPipeLine bool `json:"is_pipe_line"`

	// PipelineID is the auto-generated pipeline identifier in UUID4 format.
	PipelineID string `json:"pipeline_id,omitempty"`

	RunAtUUID string `json:"run_at_uuid,omitempty"`

	FailureReason string `json:"failure_reason,omitempty"`

	CreatedAt *time.Time `json:"created_at,omitempty"`
	// StartedAt is the UTC timestamp of the moment the job started.
	StartedAt *time.Time `json:"started_at,omitempty"`
	// CompletedAt is the UTC timestamp of the moment the job finished.
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	// LastRecyclerCreation last time the job was recycled at
	Duration time.Duration `json:"duration,omitempty"`
}
