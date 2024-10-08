package backend

import (
	"context"
	"github.com/gin-gonic/gin"
	"net"
	"strconv"
	"strings"
	"wails-amis/backend/pkg/db"
	"wails-amis/backend/pkg/response"
	"wails-amis/backend/services"
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
	db.Connect()

	b.RegisterRouter()

	b.Run()
}

func (b *Backend) Run() {
	count := 0
	portNumber, _ := strconv.Atoi(strings.TrimPrefix(b.Port, ":"))
	for {
		if count > 100 {
			panic("Server startup failed")
		}

		listener, err := net.Listen("tcp", b.Port)

		if err != nil {
			count++
			portNumber++
			b.Port = ":" + strconv.Itoa(portNumber)
			continue
		}
		_ = listener.Close()

		if err = b.Gin.Run(b.Port); err != nil {
			panic(err)
		}
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
	}
}
