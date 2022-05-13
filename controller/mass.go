package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Operations struct {
	Hosts []string `json:"host_uuids"`
}

func getOperationsBody(ctx *gin.Context) (dto *Operations, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func (base *Controller) MassOperations(c *gin.Context) {
	body, _ := getOperationsBody(c)
	//bb := new(Operations)
	var bb []string
	c.ShouldBindJSON(&bb)
	fmt.Println(bb)
	fmt.Println(body)
	//for i, a := range bb {
	//	fmt.Println(i, a)
	//}

	for i, a := range body.Hosts {
		fmt.Println(i, a)
	}
	//base.RubixProxyRequest(c)

}
