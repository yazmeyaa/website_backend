package middlewares

import (
	"net/http"
	"strings"
	"yazmeyaa_projects/data/response"
	"yazmeyaa_projects/helper"

	"github.com/gin-gonic/gin"
)

func AuthJWTMiddleware(secret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) != 2 {
			webResponse := response.Response{
				Code:   http.StatusUnauthorized,
				Status: "Unauthorized",
				Data:   nil,
			}
			ctx.JSON(http.StatusUnauthorized, webResponse)
			ctx.Abort()
			return
		}

		authToken := t[1]
		authorized, err := helper.IsAuthorized(authToken, secret)

		if !authorized {
			webResponse := response.Response{
				Code:   http.StatusUnauthorized,
				Status: "Unauthorized",
				Data:   err.Error(),
			}
			ctx.JSON(http.StatusUnauthorized, webResponse)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
