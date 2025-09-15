package router

import (
	"KanaGame/router/test"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.RedirectTrailingSlash = false
	api := r.Group("/api")
	test.RegisterApiRouter(api)

	return r
}
