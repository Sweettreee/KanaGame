package router

import (
	"KanaGame/middleware"
	"KanaGame/router/game"
	"KanaGame/router/test"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	public := r.Group("/public")

	api.Use(middleware.AuthMiddleware())
	game.RegisterGameRouter(public)
	test.RegisterApiRouter(api)

	return r
}
