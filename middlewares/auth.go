package middlewares

import (
	"github.com/SuyunovJasurbek/url_shorting/helper"
	"github.com/gin-gonic/gin"
)

func Auth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "No token found"})
		}

		param, err := helper.ValidateJWT(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
		}
		c.Set("user_id", param.UserId)
		c.Next()
	}
}
