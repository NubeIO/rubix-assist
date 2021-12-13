package schema

import (
	"github.com/NubeIO/rubix-updater/model"
	"github.com/NubeIO/rubix-updater/model/rubix"
	cmap "github.com/orcaman/concurrent-map"
	"reflect"
)

var MethodsAll = struct {
	GET    bool `json:"get"`
	POST   bool `json:"post"`
	PATCH  bool `json:"patch"`
	DELETE bool `json:"delete"`
	PUT    bool `json:"put"`
}{
	GET:    true,
	POST:   true,
	PATCH:  true,
	DELETE: true,
	PUT:    true,
}

var MethodsGetPut = struct {
	GET    bool `json:"get"`
	POST   bool `json:"post"`
	PATCH  bool `json:"patch"`
	DELETE bool `json:"delete"`
	PUT    bool `json:"put"`
}{
	GET:    true,
	POST:   false,
	PATCH:  false,
	DELETE: false,
	PUT:    true,
}

type T struct {
	Type     string `json:"type"`
	Required bool   `json:"required"`
	ReadOnly bool   `json:"read_only"`
	Min      int    `json:"min,omitempty" default:"1"`
	Max      int    `json:"max,omitempty" default:"20"`
	Default  string `json:"default,omitempty"`
}

func reflectBindings(f interface{}) cmap.ConcurrentMap {
	val := reflect.ValueOf(f).Elem()
	var obj T
	res := cmap.New()
	for i := 0; i < val.NumField(); i++ {
		//valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tag := typeField.Tag
		objType := typeField.Type.String()
		if objType == "*bool" {
			objType = "bool"
		}
		obj.Type = objType
		obj.Required = false
		obj.ReadOnly = false
		obj.Default = ""
		j := tag.Get("json")
		req := tag.Get("required")
		read := tag.Get("readonly")
		defaults := tag.Get("default")

		if req == "true" {
			obj.Required = true
		}
		if read == "true" {
			obj.ReadOnly = true
		}
		if j == "id" {
			obj.ReadOnly = true
		}
		if len(defaults) > 0 {
			obj.Default = defaults
		}
		if j != "methods" {
			if j != "-" {
				res.Set(j, obj)
			}
		}
	}
	return res
}

func GetHostSchema() interface{} {
	f := &model.Host{}
	sch := reflectBindings(f)
	sch.Set("methods", MethodsAll)
	return sch.Items()
}

func GetUserSchema() interface{} {
	f := &model.User{}
	sch := reflectBindings(f)
	sch.Set("methods", MethodsAll)
	return sch.Items()
}

func GetRubixPlatSchema() interface{} {
	f := &rubix.WiresPlat{}
	sch := reflectBindings(f)
	sch.Set("methods", MethodsGetPut)
	return sch.Items()
}

func GetRubixDiscover() interface{} {
	f := &rubix.Slaves{}
	sch := reflectBindings(f)
	sch.Set("methods", MethodsGetPut)
	return sch.Items()
}

func GetRubixSlaves() interface{} {
	f := &rubix.Slaves{}
	sch := reflectBindings(f)
	sch.Set("methods", MethodsGetPut)
	return sch.Items()
}

func GetTokenSchema() interface{} {
	f := &model.User{}
	sch := reflectBindings(f)
	sch.Set("methods", MethodsAll)
	return sch.Items()
}
