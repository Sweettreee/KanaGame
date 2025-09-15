package test

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterApiRouter(rg *gin.RouterGroup) {
	test := rg.Group("test")
	test.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{
			"message": "test",
		})
	})
}
