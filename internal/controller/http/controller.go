package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetFindToursCriteries returns the AI response to the user.
// If an error occurred during response generating or parsing it will return default response and log the error
func (h *handler) GetFindToursCriteries(ctx *gin.Context) {
	dto, err := ParseBodyDto[GetFindToursCriteriesRequest](ctx)
	if err != nil {
		// TODO: возвращать ошибку верно, учитывая то, что MustBind уже выставил все что надо
		return
	}

	response, err := h.service.GetFindToursCriteries(ctx.Request.Context(), dto.UserRequest)
	if err != nil {
		// TODO: залогировать ошибку
	}

	ctx.JSON(http.StatusOK, response)
}
