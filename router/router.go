package router

import (
	"KanaGame/middleware"
	"KanaGame/router/test"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.RedirectTrailingSlash = false
	api := r.Group("/api")

	api.Use(middleware.AuthMiddleware())
	test.RegisterApiRouter(api)

	return r
}
