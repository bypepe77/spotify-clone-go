package responses

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(code int, c *gin.Context, data interface{}) {
	c.JSON(code, Response{
		Status: true,
		Data:   data,
	})
}

func Error(code int, c *gin.Context, message string) {
	c.JSON(code, Response{
		Status:  false,
		Message: message,
	})
}
