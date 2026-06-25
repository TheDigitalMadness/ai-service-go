package http

import "github.com/gin-gonic/gin"

func (h *handler) GetFindToursCriteries(ctx *gin.Context) {
	dto, err := ParseBodyDto[GetFindToursCriteriesRequest](ctx)
	if err != nil {
		// TODO: возвращать ошибку верно, учитывая то, что MustBind уже выставил все что надо
		return
	}

}
