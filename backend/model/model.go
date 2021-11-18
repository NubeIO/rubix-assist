package model

type Host struct {
	ID            uint   `json:"id" gorm:"primarykey"`
	Name          string `json:"name"`
	IP            string `json:"ip"`
	Port          int    `json:"port"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	RubixPort     int    `json:"rubix_port"`
	RubixUsername string `json:"rubix_username"`
	RubixPassword string `json:"rubix_password"`
}

type Token struct {
	ID    uint   `json:"id" gorm:"primarykey"`
	Token string `json:"token"`
}
