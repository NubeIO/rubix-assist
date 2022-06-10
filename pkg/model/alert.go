package model

import (
	"gorm.io/datatypes"
	"time"
)

type Task struct {
	UUID          string         `json:"uuid" gorm:"primarykey"`
	Status        string         `json:"status"`
	UserUUID      string         `json:"user_uuid"`
	From          string         `json:"from"`
	HostUUID      string         `json:"host_uuid"`
	Host          string         `json:"host,omitempty"`
	Type          string         `json:"type"` //edgeapi.TaskType
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
	UUID      string         `json:"uuid" gorm:"primarykey"`
	Status    string         `json:"status"`
	Title     string         `json:"title,omitempty"`
	Type      string         `json:"type,omitempty"`
	Message   string         `json:"message,omitempty"`
	Data      datatypes.JSON `json:"data,omitempty"`
	CreatedAt time.Time      `json:"date,omitempty"`
	TaskUUID  string         `json:"task_uuid,omitempty" gorm:"TYPE:string REFERENCES tasks;"`
}
