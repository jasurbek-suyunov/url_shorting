package middlewares

import (
	"fmt"

	"github.com/SuyunovJasurbek/url_shorting/helper"
	"github.com/gin-gonic/gin"
)

func Auth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "No token found"})
		}

		token, err := helper.ValidateToken(tokenString, secret)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
		}

		fmt.Println(token)
	}
}
