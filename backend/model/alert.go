package model

import "time"

type Alert struct {
	UUID      string     `json:"uuid" gorm:"primarykey"  get:"true" delete:"true"`
	From      string     `json:"from" get:"true"`
	HostUUID  string     `json:"host_uuid" get:"true"`
	Host      string     `json:"host" get:"true"`
	AlertType string     `json:"alert_type" get:"true"`
	Count     uint       `json:"count" get:"true"`
	Date      time.Time  `json:"date" get:"true"`
	Messages  []*Message `json:"messages" gorm:"constraint:OnDelete:CASCADE"`
}

type Message struct {
	UUID      string    `json:"uuid" gorm:"primarykey"  get:"true" delete:"true"`
	Title     string    `json:"title" get:"true"`
	Message   string    `json:"message" get:"true"`
	Type      string    `json:"type" get:"true"`
	Date      time.Time `json:"date" get:"true"`
	AlertUUID string    `json:"alert_uuid" get:"true" post:"true" patch:"true" gorm:"TYPE:string REFERENCES alerts;"`
}
