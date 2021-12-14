package model

type Token struct {
	ID    string `json:"id" gorm:"primarykey"`
	Token string `json:"token"`
}
