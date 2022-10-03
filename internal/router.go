package internal

import (
	"auth/controllers/auth"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	authRouter := r.Group("/auth")
	{
		c := auth.NewController()
		authRouter.GET("/health-check", c.Healthcheck)
	}

	return r
}
