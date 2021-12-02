package schema

type Name struct {
	Type     string `json:"type" default:"string"`
	Required bool   `json:"required" default:"true"`
	Min      int    `json:"min" default:"3"`
	Max      int    `json:"max" default:"20"`
}

type Description struct {
	Type     string `json:"type" default:"string"`
	Required bool   `json:"required" default:"false"`
	Min      int    `json:"min" default:"0"`
	Max      int    `json:"max" default:"80"`
}


type Username struct {
	Type     string `json:"type" default:"string"`
	Required bool   `json:"required" default:"true"`
	Min      int    `json:"min" default:"8"`
	Max      int    `json:"max" default:"20"`
}


type Password struct {
	Type     string `json:"type" default:"string"`
	Required bool   `json:"required" default:"true"`
	Min      int    `json:"min" default:"8"`
	Max      int    `json:"max" default:"20"`
}

type IP struct {
	Type     string `json:"type" default:"string"`
	Required bool   `json:"required" default:"true"`
	Min      int    `json:"min" default:"2"`
	Max      int    `json:"max" default:"6"`
}

type Port struct {
	Type     string `json:"type" default:"string"`
	Required bool   `json:"required" default:"true"`
	Min      int    `json:"min" default:"3"`
	Max      int    `json:"max" default:"20"`
}


//type Network struct {
//	Name        NameStruct        `json:"name"`
//	Description DescriptionStruct `json:"description"`
//	PluginName  struct {
//		Type     string `json:"type" default:"string"`
//		Required bool   `json:"required" default:"true"`
//		Default  string `json:"default" default:"rubix"`
//	} `json:"plugin_name"`
//}
//
