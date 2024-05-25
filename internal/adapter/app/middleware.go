package app

import (
	"net/http"
	"strings"

	"github.com/AntonyIS/usafi-hub-cleaning-service/internal/core/ports"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type middleware struct {
	logger    ports.LoggerService
	secretKey string
}

func NewMiddleware(logger ports.LoggerService, secretKey string) *middleware {
	return &middleware{
		logger:    logger,
		secretKey: secretKey,
	}
}
func (m middleware) GinAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")
		if authorizationHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"responseCode":    http.StatusUnauthorized,
				"responseMessage": "Missing Authorization token",
			})
			c.Abort()
			return
		}

		tokenString := strings.Split(authorizationHeader, " ")[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(m.secretKey), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"responseCode":    http.StatusUnauthorized,
				"responseMessage": "Not Authorized",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
