package helpers

import (
	"final-project/params"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseStatusUnauthorized(ctx *gin.Context, additionalInfo string) {
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, params.Response{
		Status:         http.StatusUnauthorized,
		Error:          "UNAUTHORIZED",
		AdditionalInfo: additionalInfo,
	})
}

func ResponseStatusBadRequest(ctx *gin.Context, additionalInfo string) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
		Status: http.StatusBadRequest,
		Error:  "BAD REQUEST",
		AdditionalInfo: additionalInfo,
	})
}

func ResponseStatusBadRequestErrorValidation(ctx *gin.Context, additionalInfo string) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
		Status:         http.StatusBadRequest,
		Error:          "Error Validation",
		AdditionalInfo: additionalInfo,
	})
}

func ResponseStatusBadRequestCannotConvertId(ctx *gin.Context, additionalInfo string) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
		Status: http.StatusBadRequest,
		Error:  "Cannot Convert User Id",
		AdditionalInfo: additionalInfo,
	})
}
