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

var MethodsGetPostDelete = struct {
	GET    bool `json:"get"`
	POST   bool `json:"post"`
	PATCH  bool `json:"patch"`
	DELETE bool `json:"delete"`
	PUT    bool `json:"put"`
}{
	GET:    true,
	POST:   true,
	PATCH:  false,
	DELETE: true,
	PUT:    false,
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
	sch := cmap.New()
	sch.Set("fields", reflectBindings(f))
	sch.Set("METHODS", MethodsAll)
	sch.Set("METHODS", MethodsAll)
	sch.Set("HEADING", "Hosts")
	sch.Set("SUB_HEADING", "A list of hosts")
	sch.Set("HELP", "A host is an instance of the rubix system, Use the editor to add, remove, edit and delete any existing hosts")
	return sch.Items()
}

func GetUserSchema() interface{} {
	f := &model.User{}
	sch := cmap.New()
	sch.Set("fields", reflectBindings(f))
	sch.Set("METHODS", MethodsAll)
	sch.Set("HEADING", "Users")
	sch.Set("SUB_HEADING", "A list of users")
	sch.Set("HELP", "Added and remove users")
	return sch.Items()
}

func GetRubixPlatSchema() interface{} {
	f := &rubix.WiresPlat{}
	sch := cmap.New()
	sch.Set("fields", reflectBindings(f))
	sch.Set("METHODS", MethodsGetPut)
	sch.Set("HEADING", "Rubix-Details")
	sch.Set("SUB_HEADING", "site details")
	sch.Set("HELP", "Update details as required")
	return sch.Items()
}

func GetRubixDiscover() interface{} {
	f := &rubix.Slaves{}
	sch := cmap.New()
	sch.Set("fields", reflectBindings(f))
	sch.Set("METHODS", MethodsGetPut)
	sch.Set("HEADING", "Rubix-Details")
	sch.Set("SUB_HEADING", "site details")
	sch.Set("HELP", "Update details as required")
	sch.Set("API_HELP", "Update details as required")
	return sch.Items()
}

func GetRubixSlaves() interface{} {
	f := &rubix.Slaves{}
	sch := cmap.New()
	sch.Set("fields", reflectBindings(f))
	sch.Set("METHODS", MethodsGetPostDelete)
	sch.Set("HEADING", "Rubix-Details")
	sch.Set("SUB_HEADING", "site details")
	sch.Set("HELP", "Update details as required")
	sch.Set("API_HELP", "ENDPoints GET: will return a list of slaves, ADD/POST: in the body pass in the global_uuid, DELETE: to delete use the global_uuid as a parameter in the url (api/slaves/global_uuid)")
	return sch.Items()
}

func GetTokenSchema() interface{} {
	f := &model.User{}
	sch := reflectBindings(f)
	sch.Set("methods", MethodsAll)
	return sch.Items()
}
