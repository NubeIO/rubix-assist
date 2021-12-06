package model

import "time"

type Host struct {
	ID                   string    `json:"id" gorm:"primarykey"`
	Name                 string    `json:"name"  gorm:"type:varchar(255);unique;not null"`
	IP                   string    `json:"ip"`
	Port                 int       `json:"port"`
	Username             string    `json:"username"`
	Password             string    `json:"password"`
	RubixPort            int       `json:"rubix_port"`
	RubixUsername        string    `json:"rubix_username"`
	RubixPassword        string    `json:"rubix_password"`
	RubixToken           string    `json:"-"`
	RubixTokenLastUpdate time.Time `json:"-"`
	IsLocalhost          bool      `json:"is_localhost"`
}

type Token struct {
	ID    string `json:"id" gorm:"primarykey"`
	Token string `json:"token"`
}
