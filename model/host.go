package model

import "gorm.io/gorm"

type Host struct {
	gorm.Model
	IP       string `json:"ip"`
	Port     string `json:"port"`
	Username string `json:"username"`
}
