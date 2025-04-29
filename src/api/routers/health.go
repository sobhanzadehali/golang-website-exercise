package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/sobhanzadehali/golang-website-exercise.git/api/handlers"
)

func Health(r *gin.RouterGroup) {
	handler := handlers.NewHealthHandler()
	r.GET("/", handler.GetHealth)
}
