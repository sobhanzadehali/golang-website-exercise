package api

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sobhanzadehali/golang-website-exercise.git/api/routers"
	"github.com/sobhanzadehali/golang-website-exercise.git/config"
)

func InitServer() {
	cfg, err := config.GetConfig("dev")
	if err != nil {
		log.Fatal(err)
	}
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	v1 := r.Group("/api/v1")
	{
		health := v1.Group("/health")
		routers.Health(health)
	}
	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}
