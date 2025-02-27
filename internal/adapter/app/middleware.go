package app

import (
	"fmt"
	"net/http"
	"time"

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

func (m middleware) AuthorizeToken(ctx *gin.Context) {
	tokenString := ctx.GetHeader("access_token")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			m.logger.Error(fmt.Sprintf("unexpected signing method: %v", token.Header["sub"]))
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["sub"])
		}
		return []byte(m.secretKey), nil
	})

	if err != nil {
		if err != nil {
			m.logger.Error(fmt.Sprintf("Failed to verify token string : %v", err))
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"responseCode":    http.StatusUnauthorized,
				"responseMessage": "Failed to verify token string",
			})
			ctx.Abort()
			return
		}

		ctx.JSON(http.StatusUnauthorized, gin.H{
			"responseCode":    http.StatusUnauthorized,
			"responseMessage": "request not authorized",
		})
		ctx.Abort()
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"responseCode":    http.StatusUnauthorized,
				"responseMessage": "Failed to verify token string",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	} else {
		m.logger.Error("request not authorized")

		ctx.JSON(http.StatusUnauthorized, gin.H{
			"responseCode":    http.StatusUnauthorized,
			"responseMessage": "Failed to verify token string",
		})
		ctx.Abort()

		return
	}
}