package auth

import (
	"auth/internal/oauth"
	"github.com/gin-gonic/gin"
	"net/http"
)

type authorization struct {
	code string `json:"code"`
}

type Controller struct{}

func NewController() *Controller {
	return new(Controller)
}

func (*Controller) Healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": true})
}

func (*Controller) Redirect(c *gin.Context) {
	provider, err := oauth.GetAuthProvider(c.Param("provider"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, provider.Redirect())
}

func (*Controller) Callback(c *gin.Context) {
	body := &authorization{}
	err := c.ShouldBind(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	provider, err := oauth.GetAuthProvider(c.Param("provider"))
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	token, err := provider.Callback(body.code)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
