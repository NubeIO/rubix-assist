package model

import "time"

type Alert struct {
	UUID      string     `json:"uuid" gorm:"primarykey"  get:"true" post:"true" patch:"true" delete:"true"`
	From      string     `json:"from" get:"true"`
	HostUUID  string     `json:"host_uuid" get:"true" endpoint:"/api/hosts" require_key:"uuid" post_key:"host_uuid" display:"[name, IP]" field_type:"select"`
	Host      string     `json:"host" get:"true"`
	AlertType string     `json:"alert_type" get:"true"`
	Count     uint       `json:"count" get:"true"`
	Date      time.Time  `json:"date" get:"true"`
	Messages  []*Message `json:"messages" gorm:"constraint:OnDelete:CASCADE"`
}

type Message struct {
	UUID      string    `json:"uuid" gorm:"primarykey"  get:"true" post:"true" patch:"true" delete:"true"`
	Title     string    `json:"title,omitempty" get:"true"`
	Message   string    `json:"message,omitempty" get:"true"`
	Type      string    `json:"type,omitempty" get:"true"`
	Date      time.Time `json:"date,omitempty" get:"true"`
	AlertUUID string    `json:"alert_uuid,omitempty" get:"true" post:"true" patch:"true" gorm:"TYPE:string REFERENCES alerts;"`
}
