package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/utrow/golang-web-base/application/usecase"
)

type Controller interface {
	GetPing(c *gin.Context)
}

type controller struct {
	usecase usecase.Interacter
}

func NewController(it usecase.Interacter) Controller {
	return &controller{
		usecase: it,
	}
}
