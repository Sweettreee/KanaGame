package game

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterGameRouter(rg *gin.RouterGroup) {
	game := rg.Group("game")
	game.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/public/index.html", gin.H{})
	})
}
