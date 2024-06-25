package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mohammadreza-zr/golang-clean-web-api/api/routers"
	"github.com/mohammadreza-zr/golang-clean-web-api/config"
	"github.com/mohammadreza-zr/golang-clean-web-api/middlewares"
)

func InitServer() {
	cfg := config.GetConfig()
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middlewares.LimitByRequest())

	v1 := r.Group("api/v1")
	{
		health := v1.Group("/health")
		routers.Health(health)
	}

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}
