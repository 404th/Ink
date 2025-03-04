package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/404th/Ink/config"
	"github.com/404th/Ink/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// CORSMiddleware returns a middleware function that handles CORS based on environment
func CORSMiddleware(allowedOrigins []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		mode := os.Getenv("PROJECT_MODE")

		// Set common headers for all modes
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Max-Age", "3600")

		// Handle preflight OPTIONS requests first
		if c.Request.Method == "OPTIONS" {
			// In development, allow OPTIONS from anywhere
			if mode == config.ProjectModeDevelopment {
				c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
				c.AbortWithStatus(204)
				return
			}

			// In production, check origin for OPTIONS requests
			origin := c.Request.Header.Get("Origin")
			allowed := false

			for _, allowedOrigin := range allowedOrigins {
				if origin == allowedOrigin || allowedOrigin == "*" {
					allowed = true
					c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
					break
				}
			}

			if allowed {
				c.AbortWithStatus(204) // Allow the preflight
			} else {
				c.AbortWithStatus(403) // Forbidden
			}
			return
		}

		// For non-OPTIONS requests
		if mode == config.ProjectModeDevelopment {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		} else if mode == config.ProjectModeProduction {
			origin := c.Request.Header.Get("Origin")
			allowed := false

			for _, allowedOrigin := range allowedOrigins {
				if origin == allowedOrigin || allowedOrigin == "*" {
					allowed = true
					c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
					break
				}
			}

			// Block the request if origin is not allowed
			if !allowed && origin != "" {
				c.AbortWithStatus(403) // Forbidden
				return
			}
		} else {
			origin := c.Request.Header.Get("Origin")
			allowed := false

			for _, allowedOrigin := range allowedOrigins {
				if origin == allowedOrigin || allowedOrigin == "*" {
					allowed = true
					c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
					break
				}
			}

			// Block the request if origin is not allowed
			if !allowed && origin != "" {
				c.AbortWithStatus(403) // Forbidden
				return
			}
		}

		c.Next()
	}
}

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
