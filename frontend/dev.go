package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"wails-amis/backend"
)

func main() {
	app := backend.NewBackend()

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

	if err := app.Gin.Run(app.Port); err != nil {
		panic(err)
	}
}
