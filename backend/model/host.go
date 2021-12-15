package model

import "time"

type Host struct {
	ID                   string    `json:"id" gorm:"primarykey"  get:"true" delete:"true"`
	Name                 string    `json:"name"  gorm:"type:varchar(255);unique;not null" required:"true" get:"true" post:"true" patch:"true"`
	IP                   string    `json:"IP" required:"true" default:"192.168.15.10" get:"true" post:"true" patch:"true"`
	Port                 int       `json:"port" required:"true" default:"22" get:"true" post:"true" patch:"true"`
	HTTPS                *bool     `json:"HTTPS" get:"true" post:"true" patch:"true"`
	Username             string    `json:"username" required:"true" default:"admin" get:"true" post:"true" patch:"true"`
	Password             string    `json:"password" required:"true" get:"false" post:"true" patch:"true"`
	RubixPort            int       `json:"rubix_port" required:"true" default:"1660" get:"true" post:"true" patch:"true"`
	RubixUsername        string    `json:"rubix_username" required:"true" default:"admin" get:"true" post:"true" patch:"true"`
	RubixPassword        string    `json:"rubix_password" required:"true" post:"true" patch:"true"`
	IsLocalhost          bool      `json:"is_localhost" get:"true" post:"true" patch:"true"`
	RubixToken           string    `json:"-"`
	RubixTokenLastUpdate time.Time `json:"-"`
}