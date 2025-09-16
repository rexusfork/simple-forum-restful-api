package middleware

import (
	"errors"
	"net/http"
	"simple-forum/internals/configs"
	"simple-forum/pkg/internals/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	secretKey:=configs.Get().Service.SecretJWT
	return func(c *gin.Context) {
		header:= c.Request.Header.Get("Authorization")

		header = strings.TrimSpace(header)
		if header == ""{
			c.AbortWithError(http.StatusUnauthorized, errors.New("missing authorization header"))
			return 
		}

		userID, username, err:=jwt.ValidateToken(header, secretKey)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return 
		}

		c.Set("userID", userID)
		c.Set("username", username)
		c.Next()
	}
}