package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/melbahja/goph"
	"log"
)

var err error

func (base *Controller) GetPosts(c *gin.Context) {

	client, err := goph.New("pi", "192.168.15.102", goph.Password("N00BRCRC"))
	defer client.Close()

	// Execute your command.
	out, err := client.Run("pwd")

	if err != nil {
		log.Fatal(err)
	}

	// Get your output as []byte.
	fmt.Println(string(out))

	c.JSON(200, string(out))
}
