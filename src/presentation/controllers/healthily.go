package controller

import (
	usecases "backend/simple_bank/src/domain/usecases"

	"github.com/gin-gonic/gin"
)

type healthRespository interface {
	GetHealthily(ctx *gin.Context) (*usecases.GetHealthilyResult, error)
}

type Controller struct {
	hr healthRespository
}

func New(
	hr healthRespository,
) *Controller {
	return &Controller{hr}
}

func (c *Controller) GetHealthily(ctx *gin.Context) (*usecases.GetHealthilyResult, error) {
	res, err := c.hr.GetHealthily(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}
