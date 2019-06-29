package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/utrow/golang-web-base/application/usecase"
	"log"
)

type Controller interface {
	GetPing(c *gin.Context)
}

type controller struct {
	usecase usecase.Interacter
	logger  log.Logger
}

func NewController(it usecase.Interacter, logger log.Logger) Controller {
	return &controller{
		usecase: it,
		logger:  logger,
	}
}
