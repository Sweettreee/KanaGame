package router

import (
	"KanaGame/middleware"
	"KanaGame/router/test"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	game := r.Group("/game")

	api.Use(middleware.AuthMiddleware())
	game.Use()
	test.RegisterApiRouter(api)

	return r
}
