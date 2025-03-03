package api

import (
	"os"
	"time"

	"github.com/404th/Ink/api/handler"
	"github.com/404th/Ink/api/middleware"
	"github.com/404th/Ink/config"
	ginzap "github.com/gin-contrib/zap"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Run initializes and returns the Gin engine with all routes and middleware
func Run(cfg *config.Config, sugar *zap.SugaredLogger, h *handler.Handler) *gin.Engine {
	switch os.Getenv(cfg.ProjectMode) {
	case config.ProjectModeDevelopment:
		gin.SetMode(gin.DebugMode)
	case config.ProjectModeProduction:
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.TestMode)
	}

	r := gin.Default()

	r.Use(ginzap.Ginzap(sugar.Desugar(), time.RFC3339, true))

	// r.Use(CustomMiddleware()) // Custom middleware (if needed)

	r.GET("/ping", h.Ping)
	r.GET("/greeting", h.Greeting)

	// Route for refreshing tokens
	r.POST("/refresh", h.HandleRefreshJWT)

	// file
	r.POST("/file/upload/image", middleware.AuthMiddleware(cfg.AccessTokenSecret), h.UploadImageHandler)
	r.POST("/file/upload/video", middleware.AuthMiddleware(cfg.AccessTokenSecret), h.UploadVideoHandler)
	r.POST("/file/upload/file", middleware.AuthMiddleware(cfg.AccessTokenSecret), h.UploadFileHandler)

	// user
	r.POST("/api/users/login", h.LoginUser)
	r.POST("/api/users/signup", h.SignupUser)

	return r
}

// CustomMiddleware is an example of custom middleware
func CustomMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// user login validation should occur here
		c.Next()
		// Perform actions after the request is handled
	}
}
