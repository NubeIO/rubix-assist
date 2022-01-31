package schema

import (
	cmap "github.com/orcaman/concurrent-map"
	"reflect"
)

type EndPointsEdge struct {
	IP  string `json:"ip"  name:"ip-settings" endpoint:"/test" post:"true" view:"form"`
	IP2 string `json:"ip2"  name:"IP settings 2" endpoint:"/test" post:"true"  view:"table"`
}

func GetToolsEndPointsSchema() interface{} {
	f := &EndPointsEdge{}
	sch := cmap.New()
	sch.Set(endpoints, reflectBindingsEndPoint(f))
	return sch.Items()
}

type EndPointType struct {
	Name     string `json:"name"`
	Endpoint string `json:"endpoint"`
	View     string `json:"view"`
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
		get := tag.Get("get")
		post := tag.Get("post")
		patch := tag.Get("patch")
		put := tag.Get("put")
		del := tag.Get("delete")

		var obj EndPointType
		obj.Name = name
		obj.Endpoint = endpoint
		obj.View = view

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
