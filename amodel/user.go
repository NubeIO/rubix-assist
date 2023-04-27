package amodel

type User struct {
	UUID      string `json:"uuid" gorm:"primary_key"`
	Username  string `json:"username"`
	IsAdmin   bool   `json:"is_admin"`
	UserGroup bool   `json:"user_group"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	TeamID    string `json:"team"  gorm:"TYPE:string REFERENCES teams;"`
	Hash      string `json:"-"`
	UID       string `json:"-"`
	Role      string `json:"-"`
}

type NewUser struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=255"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Team struct {
	UUID  string  `json:"uuid" gorm:"primary_key"`
	Users []*User `json:"users" gorm:"constraint"`
}
