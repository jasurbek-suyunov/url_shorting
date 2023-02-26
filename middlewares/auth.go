package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "No token found"})
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				c.AbortWithStatusJSON(401, gin.H{"error": "Unexpected signing method"})
			}

			fmt.Println(token)

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return "sample", nil
		})

	}
}
