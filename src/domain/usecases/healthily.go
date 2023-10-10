package usecases

import (
	"backend/simple_bank/src/domain/model"

	"github.com/gin-gonic/gin"
)

type GetHealthily interface {
	health(ctx *gin.Context) (*GetHealthilyResult, error)
}

type GetHealthilyResult struct {
	Data model.Healthily `json:"data"`
}
