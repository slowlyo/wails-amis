package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"wails-amis/backend"
	"wails-amis/backend/pkg/db"
)

func main() {
	app := backend.NewBackend()

	db.Connect()
	app.RegisterRouter()

	app.Gin.StaticFS("/public", gin.Dir("./dist/public", false))
	app.Gin.StaticFS("/pages", gin.Dir("./dist/pages", false))
	app.Gin.NoRoute(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/api/") {
			c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
			return
		}
		c.File("./dist/index.html")
	})

	fmt.Println("Server started on: http://localhost" + app.Port)

	app.Run()
}
