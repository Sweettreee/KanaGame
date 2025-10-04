package router

import (
	"KanaGame/middleware"
	"KanaGame/router/game"
	"KanaGame/router/test"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	apiGroup := r.Group("/api")
	gameGroup := r.Group("/game")

	apiGroup.Use(middleware.AuthMiddleware())
	game.RegisterGameRouter(gameGroup)
	test.RegisterApiRouter(apiGroup)

	return r
}
