package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/utrow/golang-web-base/application/usecase/in"
	"github.com/utrow/golang-web-base/application/usecase/out"
)

func (ctr *controller) GetPing(c *gin.Context) {
	var input in.GetPing
	if c.BindQuery(&input) != nil {
		ctr.logger.Println(out.ErrInternal, input)
		c.JSON(500, out.ErrorResponse{
			Message: out.ErrInternal,
		})
		return
	}

	if input.Validate() != nil {
		c.JSON(400, out.ErrorResponse{
			Message: fmt.Sprintf(out.ErrRequire, "Text"),
		})
		return
	}

	output, err := ctr.usecase.Ping.Get(&input)
	if err != nil {
		ctr.logger.Println(out.ErrInternal, err)
		c.JSON(500, out.ErrorResponse{
			Message: out.ErrInternal,
		})
		return
	}

	c.JSON(200, output)
	return
}
