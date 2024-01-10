package middleware

import (
	"fmt"
	"log"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Test() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Athorization header is missing"})
			return
		}
		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid authorization header"})
			return
		}
		token, err := jwt.Parse(authParts[1], func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
			}
			return []byte("not_so_scret"), nil
		})

		if err != nil || !token.Valid {
			log.Println(err)
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid JWT"})
			return
		}

		c.Next()
	}
}
