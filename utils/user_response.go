package utils

import (
	"net/http"
	"todo-api/models/dto"

	"github.com/gin-gonic/gin"
)

// SendSingleResponse mengirimkan respon tunggal
func SendSingleResponse(ctx *gin.Context, message string, data any, code int) {
	ctx.JSON(http.StatusOK, dto.SingleResponse{
		Status: dto.Status{
			Code:    code,
			Message: message,
		},
		Data: data,
	})
}

// SendPagingResponse mengirimkan respon paging
func SendPagingResponse(ctx *gin.Context, message string, data []any, paging dto.Paging, code int) {
	ctx.JSON(http.StatusOK, dto.PagingResponse{
		Status: dto.Status{
			Code:    code,
			Message: message,
		},
		Data:   data,
		Paging: paging,
	})
}

// SendErrorResponse mengirimkan respon error
func SendErrorResponse(ctx *gin.Context, message string, code int) {
	ctx.JSON(code, dto.SingleResponse{
		Status: dto.Status{
			Code:    code,
			Message: message,
		},
	})
}
