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

	var m []model.Host
	if err := base.DB.Find(&m).Error; err != nil {
		logger.Errorf("GetPost error: %v", err)
		c.JSON(200, err)
	} else {
		c.JSON(200, m)
	}

}

func (base *Controller) GetHost(c *gin.Context) {
	m := new(model.Host)
	id := c.Params.ByName("id")
	if err := base.DB.Where("id = ? ", id).First(&m).Error; err != nil {
		logger.Errorf("GetHost error: %v", err)
		reposeHandler(m, err, c)
	} else {
		reposeHandler(m, nil, c)
	}
}

func (base *Controller) CreateHost(c *gin.Context) {
	m := new(model.Host)
	err := c.ShouldBindJSON(&m)
	if err := base.DB.Create(&m).Error; err != nil {
		logger.Errorf("CreateHost error: %v", err)
	}
	reposeHandler(m, err, c)
}

func getHostBody(ctx *gin.Context) (dto *model.Host, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func (base *Controller) UpdateHost(c *gin.Context) {
	m := new(model.Host)
	id := c.Params.ByName("id")
	body, _ := getHostBody(c)
	query := base.DB.Where("id = ?", id).Find(&m)
	query = base.DB.Model(&m).Updates(body)
	if query.Error != nil {
		reposeHandler(m, query.Error, c)
	} else {
		reposeHandler(m, nil, c)
	}
}

func (base *Controller) DeleteHost(c *gin.Context) {
	m := new(model.Host)
	id := c.Params.ByName("id")
	if err := base.DB.Where("id = ? ", id).Delete(&m).Error; err != nil {
		logger.Errorf("GetHost error: %v", err)
		reposeHandler(m, err, c)
	} else {
		reposeHandler(m, nil, c)
	}
}
