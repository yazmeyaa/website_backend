package helper

import (
	"net/http"
	"yazmeyaa_projects/data/response"

	"github.com/gin-gonic/gin"
)

func HandleHTTPError(ctx *gin.Context, status int, err error) {
	ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
		Error: err.Error(),
	})
}
