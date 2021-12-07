package schema

import (
	"github.com/NubeIO/rubix-updater/model/schema/defaults"
)

var MethodsAll = struct {
	GET    bool `json:"get"`
	POST   bool `json:"post"`
	PATCH  bool `json:"patch"`
	DELETE bool `json:"delete"`
}{
	GET:    true,
	POST:   true,
	PATCH:  true,
	DELETE: true,
}

type StringRequired struct {
	Type     string `json:"type" default:"string"`
	Required bool   `json:"required" default:"true"`
	Min      int    `json:"min" default:"1"`
	Max      int    `json:"max" default:"30"`
}

type StringNotRequired struct {
	Type     string `json:"type" default:"string"`
	Required bool   `json:"required" default:"false"`
	Min      int    `json:"min" default:"1"`
	Max      int    `json:"max" default:"30"`
}

type BoolRequired struct {
	Type     string `json:"type" default:"boolean"`
	Required bool   `json:"required" default:"true"`
}

type BoolNotRequired struct {
	Type     string `json:"type" default:"boolean"`
	Required bool   `json:"required" default:"false"`
}

type ID struct {
	Type     string `json:"type" default:"string"`
	Required bool   `json:"required" default:"false"`
	ReadOnly bool   `json:"read_only" default:"true"`
}

type Host struct {
	Methods       interface{}       `json:"methods"`
	ID            ID                `json:"id"`
	Name          StringRequired    `json:"name"`
	Description   StringNotRequired `json:"description"`
	Username      StringRequired    `json:"username"`
	Password      StringRequired    `json:"password"`
	IP            IP                `json:"ip"`
	Port          Port              `json:"port"`
	RubixUsername StringRequired    `json:"rubix_username"`
	RubixPassword StringRequired    `json:"rubix_password"`
	RubixPort     Port              `json:"rubix_port"`
}

type User struct {
	Methods   interface{}       `json:"methods"`
	ID        ID                `json:"id"`
	Name      StringRequired    `json:"name"`
	Username  StringRequired    `json:"username"`
	Password  StringRequired    `json:"password"`
	Email     StringRequired    `json:"email"`
	UserGroup StringNotRequired `json:"user_group"`
	IsAdmin   BoolNotRequired   `json:"is_admin"`
}

func GetHostSchema() *Host {
	s := &Host{
		Methods: MethodsAll,
	}
	defaults.Set(s)
	return s
}

func GetTokenSchema() *Host {
	s := &Host{
		Methods: MethodsAll,
	}
	defaults.Set(s)
	return s
}

func GetUserSchema() *User {
	s := &User{
		Methods: MethodsAll,
	}
	defaults.Set(s)
	return s
}
