package middleware

import (
	"net/http"
	"strings"

	"github.com/404th/Ink/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware(accessSecretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		tokenStringArr := strings.Split(tokenString, " ")
		if len(tokenStringArr) != 2 || tokenStringArr[0] != "Bearer" {
			var errResp model.Response
			errResp.Success = false
			errResp.Message = "Avtorizatsiyadan o'ting"
			errResp.Data = nil
			c.AbortWithStatusJSON(http.StatusUnauthorized, errResp)
			return
		}

		// Parse the access token
		token, err := jwt.Parse(tokenStringArr[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrAbortHandler
			}
			return []byte(accessSecretKey), nil
		})

		// Check if token is valid
		if err != nil || !token.Valid {
			var errResp model.Response
			errResp.Message = "Tizimga qayta kiring"
			errResp.Data = nil
			c.AbortWithStatusJSON(http.StatusUnauthorized, errResp)
			return
		}

		// Set claims in the context
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("_x_data", claims)
		} else {
			var errResp model.Response
			errResp.Message = "Tizimga qayta kiring"
			errResp.Data = nil
			c.AbortWithStatusJSON(http.StatusUnauthorized, errResp)
			return
		}

		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := c.MustGet("_x_data").(jwt.MapClaims)
		role := claims["role"].(string)

		if role != "admin" {
			var errResp model.Response
			errResp.Message = "Ruxsat etilmagan"
			errResp.Data = nil
			c.AbortWithStatusJSON(http.StatusUnauthorized, errResp)
			return
		}

		c.Next()
	}
}
