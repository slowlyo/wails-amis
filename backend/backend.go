package backend

import (
	"context"
	"github.com/gin-gonic/gin"
	"wails-amis/backend/services"
	"wails-amis/backend/utils/response"
)

type Backend struct {
	Gin  *gin.Engine
	Port string
}

func NewBackend() *Backend {
	return &Backend{
		Gin:  gin.New(),
		Port: ":32001",
	}
}

func (b *Backend) HttpHandler() *gin.Engine {
	return b.Gin
}

func (b *Backend) Startup(ctx context.Context) {
	b.RegisterRouter()

	if err := b.Gin.Run(b.Port); err != nil {
		panic(err)
	}
}

func (b *Backend) RegisterRouter() {
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
