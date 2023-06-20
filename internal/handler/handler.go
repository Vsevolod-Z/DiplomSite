package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	main := router.Group("/")
	{
		main.GET("main", h.main)
		app := main.Group("app")
		{
			app.GET("", h.game)
		}

	}

	return router
}
