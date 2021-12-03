package model

type User struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Username  string `json:"username"`
	IsAdmin   bool   `json:"is_admin"`
	UserGroup bool   `json:"user_group"`
	Hash      string `json:"-"`
	Email     string `json:"email"`
	UID       string `json:"uid"`
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
