package internal

import (
	"auth/controllers/auth"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	v1 := r.Group("/v1")
	v1Auth := v1.Group("/auth")
	{
		c := auth.NewController()
		v1Auth.GET("/health-check", c.Healthcheck)
	}

	return r
}
