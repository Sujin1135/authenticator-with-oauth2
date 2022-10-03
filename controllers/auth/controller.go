package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct{}

func NewController() *Controller {
	return new(Controller)
}

func (*Controller) Healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": true})
}
