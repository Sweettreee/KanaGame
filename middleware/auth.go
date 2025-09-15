package middleware

import (
	jwt "KanaGame/jwtutils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("Login", false)
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.Next()
			return
		}

		token := parts[1]
		if uid, err := jwt.VerifyToken(token); err != nil {
			if uid != 0 {
				c.Set("Login", true)
				c.Set("Uid", uid)
				c.Next()
				return
			}
		}
		c.Next()
	}

}
