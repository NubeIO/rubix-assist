package model

type Token struct {
	UUID    string `json:"uuid" gorm:"primarykey"`
	Token string `json:"token"`
}
