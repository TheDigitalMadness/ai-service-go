package http

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func ParseBodyDto[dtoType any](ctx *gin.Context) (*dtoType, error) {
	var dto dtoType
	if err := ctx.MustBindWith(&dto, binding.JSON); err != nil {
		return nil, err
	}
	return &dto, nil
}
