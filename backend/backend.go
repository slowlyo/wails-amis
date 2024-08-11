package backend

import (
	"context"
	"github.com/gin-gonic/gin"
	"wails-amis/backend/services"
	"wails-amis/backend/utils/response"
)

type Backend struct {
	Gin *gin.Engine
}

func NewBackend() *Backend {
	return &Backend{
		Gin: gin.New(),
	}
}

func (b *Backend) HttpHandler() *gin.Engine {
	return b.Gin
}

func (b *Backend) Startup(ctx context.Context) {
	b.registerRouter()

	if err := b.Gin.Run(":32001"); err != nil {
		panic(err)
	}
}

func (b *Backend) registerRouter() {
	api := b.Gin.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			response.Ok(c, "pong")
		})

		hello := services.HelloService{}
		api.GET("/hello", hello.Index)
		api.Group("/api")
	}
}
