package http

import "github.com/gin-gonic/gin"

type HttpHandler interface {
	GetFindToursCriteries(ctx *gin.Context)
}

func NewRouter(handler HttpHandler) *gin.Engine {
	router := gin.Default()

	router.POST("/get-find-tours-criteries", handler.GetFindToursCriteries)

	return router
}
