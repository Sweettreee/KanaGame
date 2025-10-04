package router

import (
	"KanaGame/middleware"
	"KanaGame/router/game"
	"KanaGame/router/test"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("public/templates/*")
	r.Static("/static", "public/static")

	r.GET("/", func(c *gin.Context) {
		c.Redirect(301, "/public/game")
	})

	api := r.Group("/api")
	public := r.Group("/public")

	r.Use(middleware.AuthMiddleware())
	game.RegisterGameRouter(public)
	test.RegisterApiRouter(api)

	return r
}
