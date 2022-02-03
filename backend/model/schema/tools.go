package schema

import (
	cmap "github.com/orcaman/concurrent-map"
	"reflect"
)

type EPEdgeIP struct {
	ArchType string `json:"alerts"  name:"alerts" help:"get alerts" endpoint:"/alerts" get:"true"  get:"true" post:"true" post:"patch" view:"table"`
	//IP     string `json:"ip"   name:"ip-settings"  help:"set the ip on the edge-28 to a fixed ip address"  endpoint:"/tools/edge/ip" post:"true" view:"form"`
	//IpDHCP string `json:"ip_dhcp"  name:"ip-dbcp"   help:"set the ip on the edge-28 to auto dhcp"  endpoint:"/tools/edge/ip/dhcp" post:"true" view:"form"`
}

type EPSystem struct {
	ArchType string `json:"users"  name:"users" help:"get users" endpoint:"/users" get:"true" view:"table"`
	NODEJS   string `json:"nodejs"  name:"nodejs"  help:"get nodejs version" endpoint:"/tools/nodejs" get:"true" view:"form"`
}

func GetToolsEndPointsSchema() interface{} {
	e1 := &EPEdgeIP{}
	e2 := &EPSystem{}
	sch := cmap.New()
	sch.Set("alerts", reflectBindingsEndPoint(e1))
	sch.Set("programs", reflectBindingsEndPoint(e2))
	return sch.Items()
}

type EndPointType struct {
	Name     string `json:"name"`
	Endpoint string `json:"endpoint"`
	View     string `json:"view"`
	Help     string `json:"help"`
	Get      bool   `json:"get"`
	Post     bool   `json:"post"`
	Patch    bool   `json:"patch"`
	Put      bool   `json:"put"`
	Delete   bool   `json:"delete"`
}

func reflectBindingsEndPoint(f interface{}) cmap.ConcurrentMap {
	val := reflect.ValueOf(f).Elem()
	res := cmap.New()
	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i)
		tag := typeField.Tag
		j := tag.Get("json")
		name := tag.Get("name")
		endpoint := tag.Get("endpoint")
		view := tag.Get("view")
		help := tag.Get("help")
		get := tag.Get("get")
		post := tag.Get("post")
		patch := tag.Get("patch")
		put := tag.Get("put")
		del := tag.Get("delete")

		var obj EndPointType
		obj.Name = name
		obj.Endpoint = endpoint
		obj.View = view
		obj.Help = help

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
		if j != "methods" {
			if j != "-" {
				res.Set(j, obj)
			}
		}

	}
	return res
}
