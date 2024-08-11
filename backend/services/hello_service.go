package services

import (
	"github.com/gin-gonic/gin"
	"wails-amis/backend/utils/response"
)

type HelloService struct {
}

func (h *HelloService) Index(c *gin.Context) {
	response.Success(c, response.R{
		Data: gin.H{
			"type": "page",
			"body": "Hello Wails amis",
		},
	})
}
