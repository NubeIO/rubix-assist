package model

import "time"

type Alert struct {
	ID       string     `json:"id" gorm:"primarykey"  get:"true" delete:"true"`
	From     string     `json:"from" get:"true"`
	HostID   string     `json:"host_id" get:"true"`
	Host     string     `json:"host" get:"true"`
	Count    uint       `json:"count" get:"true"`
	Messages []*Message `json:"messages" gorm:"constraint"`
}

type Message struct {
	ID      string    `json:"id" gorm:"primarykey"  get:"true" delete:"true"`
	Title   string    `json:"title" get:"true"`
	Message string    `json:"message" get:"true"`
	Type    string    `json:"type" get:"true"`
	Date    time.Time `json:"date" get:"true"`
	AlertID string    `json:"alerts" get:"true" post:"true" patch:"true" gorm:"TYPE:string REFERENCES alerts;"`
}
