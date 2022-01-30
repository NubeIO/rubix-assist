package schema

import (
	"github.com/NubeIO/rubix-updater/model"
	cmap "github.com/orcaman/concurrent-map"
)

var edgeHelp = `line 1
 line 2
 line 3`

func GetEdge28IPSchema() interface{} {
	f := &model.Message{}
	sch := cmap.New()
	sch.Set(fields, reflectBindings(f))
	sch.Set(methods, MethodsPost)
	sch.Set(heading, "Edge-28 Network Settings")
	sch.Set(subHeading, "A list of Messages")
	sch.Set(help, edgeHelp)
	return sch.Items()
}
