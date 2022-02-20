package schema

import (
	cmap "github.com/orcaman/concurrent-map"
	"reflect"
)

type AlertsTools struct {
	EndPoint  string `json:"alerts"  name:"alerts" help:"get alerts" endpoint:"/alerts" get:"true"  view:"table"`
	TableLink string `json:"table_link"  name:"messages" help:"get messages by alert id" get:"true" field:"name" endpoint:"/alerts/:uuid" link_type:"by_uuid" link_flied:"uuid" schema:"messages/schema"  view:"table"`
	//Schema   interface{} `json:"schema"`
	//IP     string `json:"ip"   name:"ip-settings"  help:"set the ip on the edge-28 to a fixed ip address"  endpoint:"/tools/edge/ip" post:"true" view:"form"`
	//IpDHCP string `json:"ip_dhcp"  name:"ip-dbcp"   help:"set the ip on the edge-28 to auto dhcp"  endpoint:"/tools/edge/ip/dhcp" post:"true" view:"form"`
}

//type EPEdgeIP struct {
//	ArchType string `json:"alert"  name:"alerts" help:"get alerts" endpoint:"/alerts" get:"true"  view:"table"`
//	//IP     string `json:"ip"   name:"ip-settings"  help:"set the ip on the edge-28 to a fixed ip address"  endpoint:"/tools/edge/ip" post:"true" view:"form"`
//	//IpDHCP string `json:"ip_dhcp"  name:"ip-dbcp"   help:"set the ip on the edge-28 to auto dhcp"  endpoint:"/tools/edge/ip/dhcp" post:"true" view:"form"`
//}

type EPSystem struct {
	ArchType string `json:"users"  name:"users" help:"get users" endpoint:"/users" get:"true" view:"table"`
	NODEJS   string `json:"nodejs"  name:"nodejs"  help:"get nodejs version" endpoint:"/tools/nodejs" get:"true" view:"form"`
}

func GetToolsEndPointsSchema() interface{} {
	e1 := &AlertsTools{}
	o := map[string]map[string]interface{}{
		"alerts": {
			"name":     "alerts",
			"schema":   GetAlertSchema(),
			"endpoint": reflectBindingsEndPoint(e1),
		},
		"alerts2": {
			"name":     "alerts two",
			"schema":   GetAlertSchema(),
			"endpoint": reflectBindingsEndPoint(e1),
		},
	}

	//o := map[string]map[string]interface{}{
	//	"alerts": {
	//		"name":     "alerts",
	//		"schema":   GetAlertSchema(),
	//		"endpoint": reflectBindingsEndPoint(e1),
	//		"child": map[string]interface{}{
	//			"name":     "name 1",
	//			"schema":   GetAlertSchema(),
	//			"endpoint": reflectBindingsEndPoint(e1),
	//		},
	//	},
	//	"alerts2": {
	//		"name":     "alerts two",
	//		"schema":   GetAlertSchema(),
	//		"endpoint": reflectBindingsEndPoint(e1),
	//	},
	//}

	return o
}

type EndPointType struct {
	Name      string `json:"name"`
	Endpoint  string `json:"endpoint"`
	View      string `json:"view"`
	Field     string `json:"field"`
	LinkType  string `json:"link_type"`
	LinkFlied string `json:"link_flied"`
	Schema    string `json:"schema"`
	Help      string `json:"help"`
	Get       bool   `json:"get"`
	Post      bool   `json:"post"`
	Patch     bool   `json:"patch"`
	Put       bool   `json:"put"`
	Delete    bool   `json:"delete"`
}

func reflectBindingsEndPoint(f interface{}) interface{} {
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

		field := tag.Get("field")
		linkType := tag.Get("link_type")
		linkFlied := tag.Get("link_flied")
		schema := tag.Get("schema")

		var obj EndPointType
		obj.Name = name
		obj.Endpoint = endpoint
		obj.View = view
		obj.Help = help
		obj.Field = field
		obj.LinkType = linkType
		obj.LinkFlied = linkFlied
		obj.Schema = schema

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
