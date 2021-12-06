package schema

import "github.com/NubeIO/rubix-updater/model/schema/defaults"

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

type Host struct {
	Methods     interface{} `json:"methods"`
	Name        Name        `json:"name"`
	Description Description `json:"description"`
	Username    Password    `json:"username"`
	Password    Password    `json:"password"`
	IP          IP          `json:"ip"`
	Port        Port        `json:"port"`
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

func GetUserSchema() *Host {
	s := &Host{
		Methods: MethodsAll,
	}
	defaults.Set(s)
	return s
}
