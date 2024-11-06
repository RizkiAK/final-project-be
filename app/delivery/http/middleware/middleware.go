package middleware

import (
	"blog-mandalika/domain"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/Yureka-Teknologi-Cipta/yureka/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type AppMiddleware struct {
}

func (m *AppMiddleware) CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Header.Set("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func (m *AppMiddleware) AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		hAuth := c.GetHeader("Authorization")
		if hAuth == "" {
			response := response.Error(http.StatusBadRequest, "Unauthorized: Header authorization is required")
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		splitToken := strings.Split(hAuth, "Bearer ")
		if len(splitToken) != 2 {
			response := response.Error(http.StatusBadRequest, "Unauthorized: Token is invalid")
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		// get token without 'Bearer '
		tokenString := splitToken[1]
		// validating token
		token, err := jwt.ParseWithClaims(tokenString, &domain.JWTClaimAdmin{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		// check validity token
		if !token.Valid {
			msg := err.Error()
			if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
				msg = "Unauthorized: Token signature invalid"
			}
			response := response.Error(http.StatusBadRequest, msg)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		claims, tokenOK := token.Claims.(*domain.JWTClaimAdmin)
		if !tokenOK {
			response := response.Error(http.StatusBadRequest, "Unauthorized: Token data not valid")
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		if !token.Valid {
			response := response.Error(http.StatusBadRequest, err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		c.Set("id", claims.UserID)
		c.Next()
	}
}

// InitMiddleware initialize the middleware
func InitMiddleware() *AppMiddleware {
	return &AppMiddleware{}
}
