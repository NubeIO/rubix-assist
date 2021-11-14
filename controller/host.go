package controller

import (
	"github.com/NubeIO/rubix-updater/model"
	"github.com/NubeIO/rubix-updater/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Message struct {
	Message string `json:"message"`
}

func reposeHandler(body interface{}, err error, ctx *gin.Context) {
	if err != nil {
		if body == nil {
			ctx.JSON(404, Message{Message: "unknown error"})
		} else {
			ctx.JSON(404, Message{Message: err.Error()})
		}
	} else {
		ctx.JSON(200, body)
	}
}

func (base *Controller) GetPosts(c *gin.Context) {

	//client, err := goph.New("pi", "192.168.15.102", goph.Password("N00BRCRC"))
	//defer client.Close()
	//
	//// Execute your command.
	//out, err := client.Run("pwd")
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// Get your output as []byte.
	//fmt.Println(string(out))


	m := new(model.Host)
	if err := base.DB.Where("id = ? ", "id").Preload("Tags").First(&m).Error; err != nil {
		logger.Errorf("GetPost error: %v", err)
		c.JSON(200, err)
	}
	c.JSON(200, m)
}
func (base *Controller) CreateHost(c *gin.Context) {
	m := new(model.Host)
	err := c.ShouldBindJSON(&m)
	if err :=  base.DB.Save(&m).Error; err != nil {
		logger.Errorf("CreateHost error: %v", err)
	}
	reposeHandler(m, err, c)
}
