package router

import (
	"github.com/gin-gonic/gin"
	"github.com/utrow/golang-web-base/interface/controller"
)

func NewRouter(r *gin.Engine, c controller.Controller) *gin.Engine {
	r.GET("/ping", c.GetPing)

	return r
}
