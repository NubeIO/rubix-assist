package amodel

type Token struct {
	UUID  string `json:"uuid" gorm:"primarykey"`
	Token string `json:"token"`
}
