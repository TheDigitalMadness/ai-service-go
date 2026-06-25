package http

import "github.com/gin-gonic/gin"

type HttpHandler interface {
	GetFindToursCriteries(ctx *gin.Context)
}
