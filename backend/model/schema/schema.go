package schema

import (
	"fmt"
	"github.com/NubeIO/rubix-updater/model"
	"github.com/NubeIO/rubix-updater/model/rubix"
	cmap "github.com/orcaman/concurrent-map"
	"reflect"
	"strings"
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

var MethodsPost = struct {
	GET    bool `json:"get"`
	POST   bool `json:"post"`
	PATCH  bool `json:"patch"`
	DELETE bool `json:"delete"`
	PUT    bool `json:"put"`
}{
	GET:    false,
	POST:   true,
	PATCH:  false,
	DELETE: false,
	PUT:    false,
}

type Actions struct {
	VIEW   string `json:"view"`
	ADD    bool   `json:"add"`
	EDIT   bool   `json:"edit"`
	DELETE int    `json:"delete"`
}

type T struct {
	Type     string `json:"type"`
	Required bool   `json:"required"`
	Min      int    `json:"min,omitempty" default:"1"`
	Max      int    `json:"max,omitempty" default:"20"`
	Default  string `json:"default,omitempty"`
	Get      bool   `json:"get"`
	Post     bool   `json:"post"`
	Patch    bool   `json:"patch"`
	Put      bool   `json:"put"`
	Delete   bool   `json:"delete"`
}

func reflectBindings(f interface{}) cmap.ConcurrentMap {
	val := reflect.ValueOf(f).Elem()
	res := cmap.New()
	for i := 0; i < val.NumField(); i++ {
		//valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tag := typeField.Tag
		objType := typeField.Type.String()
		if objType == "*bool" {
			objType = "bool"
		}
		if strings.Contains(objType, "[]*") {
			objType = "array"
		}
		if objType == "time.Time" {
			objType = "date"
		}
		var obj T
		obj.Type = objType
		obj.Default = ""
		j := tag.Get("json")
		req := tag.Get("required")
		defaults := tag.Get("default")
		get := tag.Get("get")
		post := tag.Get("post")
		patch := tag.Get("patch")
		put := tag.Get("put")
		del := tag.Get("delete")

		name := tag.Get("name")
		endpoint := tag.Get("endpoint")

		fmt.Println(name, endpoint)

		if req == "true" {
			obj.Required = true
		}
		if get == "true" {
			obj.Get = true
		}
		if post == "true" {
			obj.Post = true
		}
		if patch == "true" {
			obj.Patch = true
		}
		if put == "true" {
			obj.Put = true
		}
		if del == "true" {
			obj.Delete = true
		}
		if req == "true" {
			obj.Required = true
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

const (
	fields     = "fields"
	methods    = "methods"
	heading    = "heading"
	subHeading = "sub_heading"
	help       = "help"
	apiHelp    = "api_help"
	view       = "view"
	endpoints  = "endpoints"
)

func GetHostSchema() interface{} {
	f := &model.Host{}
	sch := cmap.New()
	sch.Set(fields, reflectBindings(f))
	sch.Set(methods, MethodsAll)
	sch.Set(heading, "Hosts")
	sch.Set(subHeading, "A list of hosts")
	sch.Set(help, "A host is an instance of the rubix system, Use the editor to add, remove, edit and delete any existing hosts")
	sch.Set(view, "table")
	return sch.Items()
}

func GetDeviceInfoSchema() interface{} {
	f := &model.Host{}
	sch := cmap.New()
	sch.Set(fields, reflectBindings(f))
	sch.Set(methods, MethodsAll)
	sch.Set(heading, "Hosts")
	sch.Set(subHeading, "A list of hosts")
	sch.Set(help, "A host is an instance of the rubix system, Use the editor to add, remove, edit and delete any existing hosts")
	return sch.Items()
}

func GetUserSchema() interface{} {
	f := &model.User{}
	sch := cmap.New()
	sch.Set(fields, reflectBindings(f))
	sch.Set(methods, MethodsAll)
	sch.Set(heading, "Users")
	sch.Set(subHeading, "A list of users")
	sch.Set(help, "Added and remove users")
	return sch.Items()
}

func GetTeamSchema() interface{} {
	f := &model.Team{}
	sch := cmap.New()
	sch.Set(fields, reflectBindings(f))
	sch.Set(methods, MethodsAll)
	sch.Set(heading, "Teams")
	sch.Set(subHeading, "A list of teams")
	sch.Set(help, "Added and remove teams")
	return sch.Items()
}

func GetAlertSchema() interface{} {
	f := &model.Alert{}
	sch := cmap.New()
	sch.Set(fields, reflectBindings(f))
	sch.Set(methods, MethodsAll)
	sch.Set(heading, "Alerts")
	sch.Set(subHeading, "A list of Alerts")
	sch.Set(help, "Added and remove Alerts")
	return sch.Items()
}

func GetMessageSchema() interface{} {
	f := &model.Message{}
	sch := cmap.New()
	sch.Set(fields, reflectBindings(f))
	sch.Set(methods, MethodsAll)
	sch.Set(heading, "Messages")
	sch.Set(subHeading, "A list of Messages")
	sch.Set(help, "Added and remove Messages")
	return sch.Items()
}

func GetRubixPlatSchema() interface{} {
	f := &rubix.WiresPlat{}
	sch := cmap.New()
	sch.Set(fields, reflectBindings(f))
	sch.Set(methods, MethodsGetPut)
	sch.Set(heading, "Rubix-Details")
	sch.Set(subHeading, "site details")
	sch.Set(help, "Update details as required")
	return sch.Items()
}

func GetRubixDiscover() interface{} {
	f := &rubix.Slaves{}
	sch := cmap.New()
	sch.Set(fields, reflectBindings(f))
	sch.Set(methods, MethodsGetPut)
	sch.Set(heading, "Rubix-Details")
	sch.Set(subHeading, "site details")
	sch.Set(help, "Update details as required")
	return sch.Items()
}

func GetRubixSlaves() interface{} {
	f := &rubix.Slaves{}
	sch := cmap.New()
	sch.Set(fields, reflectBindings(f))
	sch.Set(methods, MethodsGetPostDelete)
	sch.Set(heading, "Rubix-Details")
	sch.Set(subHeading, "site details")
	sch.Set(help, "Update details as required")
	sch.Set(apiHelp, "ENDPoints GET: will return a list of slaves, ADD/POST: in the body pass in the global_uuid, DELETE: to delete use the global_uuid as a parameter in the url (api/slaves/global_uuid)")
	return sch.Items()
}

func GetTokenSchema() interface{} {
	f := &model.User{}
	sch := reflectBindings(f)
	sch.Set("methods", MethodsAll)
	return sch.Items()
}
