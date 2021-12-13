package model

import "time"

type Host struct {
	ID                   string    `json:"id" gorm:"primarykey" readonly:"true"`
	Name                 string    `json:"name"  gorm:"type:varchar(255);unique;not null" required:"true"`
	IP                   string    `json:"ip" required:"true" default:"192.168.15.10"`
	Port                 int       `json:"port" required:"true" default:"22"`
	HTTPS                *bool     `json:"https"`
	Username             string    `json:"username" required:"true" default:"admin"`
	Password             string    `json:"password" required:"true"`
	RubixPort            int       `json:"rubix_port" required:"1660"`
	RubixUsername        string    `json:"rubix_username" required:"true" default:"admin"`
	RubixPassword        string    `json:"rubix_password" required:"true"`
	RubixToken           string    `json:"-"`
	RubixTokenLastUpdate time.Time `json:"-"`
	IsLocalhost          bool      `json:"is_localhost"`
}

type Token struct {
	ID    string `json:"id" gorm:"primarykey"`
	Token string `json:"token"`
}
