package response

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Response struct
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// WithData response request
func WithData(c *gin.Context, httpCode, responseCode int, data interface{}) {
	response := Response{
		Code:    responseCode,
		Message: GetMsg(responseCode),
		Data:    data,
	}
	c.JSON(httpCode, response)
	if mode := gin.Mode(); mode == gin.DebugMode {
		switch data.(type) {
		case error:
			log.Panicf("[error] %+v", data)
		}
	}
}
