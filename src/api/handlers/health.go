package handlers

import "github.com/gin-gonic/gin"

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) GetHealth(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Healthy",
	})

}
