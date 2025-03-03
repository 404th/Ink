package handler

import (
	"fmt"
	"net/http"

	"github.com/404th/Ink/config"
	"github.com/404th/Ink/internal/service"
	"github.com/404th/Ink/internal/tigres"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler struct {
	cfg     *config.Config
	sugar   *zap.SugaredLogger
	service service.ServiceI
	tg      tigres.TigresI
}

func NewHandler(cfg *config.Config, sugar *zap.SugaredLogger, service service.ServiceI, tg tigres.TigresI) *Handler {
	return &Handler{
		cfg:     cfg,
		sugar:   sugar,
		service: service,
		tg:      tg,
	}
}

func (h *Handler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "Pong")
}

func (h *Handler) Greeting(c *gin.Context) {
	c.JSON(http.StatusOK, fmt.Sprintf("Heeey, what's up %s", c.Query("name")))
}
