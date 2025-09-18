package test

import (
	mysql "KanaGame/mysqlclient"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterApiRouter(rg *gin.RouterGroup) {
	test := rg.Group("test")
	test.GET("/", func(c *gin.Context) {
		mysql.GetMysqlConnection()
		c.JSON(http.StatusCreated, gin.H{
			"message": "test",
		})
	})
}
