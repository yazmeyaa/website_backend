package controller

import (
	"context"
	"fmt"
	"net/http"
	"yazmeyaa_projects/data/request"
	"yazmeyaa_projects/data/response"
	"yazmeyaa_projects/service"

	"github.com/gin-gonic/gin"
)

type CorsController struct {
	corsService service.CorsService
}

func NewCorsController(cs service.CorsService) *CorsController {
	return &CorsController{
		corsService: cs,
	}
}

func (cc *CorsController) AddOrigin(ctx *gin.Context) {
	origin := request.AddCorsOriginRequest{}
	if err := ctx.ShouldBindJSON(&origin); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Error",
		})
		return
	}

	record, err := cc.corsService.AddOrigin(ctx, origin.Origin, origin.AllowedMethods, origin.AllowedHeaders)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Error",
		})
		return
	}

	ctx.JSON(http.StatusOK, response.AddCorsOriginResponse{
		Origin:         origin.Origin,
		AllowedMethods: record.AllowedMethods,
		AllowedHeaders: record.AllowedHeaders,
	})
}

func (cc *CorsController) EnableOrigin(ctx *gin.Context) {
	panic("unimplemented")
}

func (cc *CorsController) DisableOrigin(ctx *gin.Context) {
	panic("unimplemented")
}

func (cc *CorsController) RemoveOrigin(ctx *gin.Context) {
	origin := ctx.Param("origin")
	if origin == "" {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "Empty origin param",
		})
		return
	}

	if err := cc.corsService.RemoveOrigin(context.Background(), origin); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "Empty origin param",
		})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (cc *CorsController) GetAllOrigins(ctx *gin.Context) {
	origins, err := cc.corsService.GetAllRecords(context.Background())
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Error",
		})
		return
	}
	records := make([]response.AddCorsOriginResponse, len(origins))

	for idx, val := range origins {
		records[idx] = response.AddCorsOriginResponse{
			Origin:         val.Origin,
			AllowedMethods: val.AllowedMethods,
			AllowedHeaders: val.AllowedHeaders,
			OriginAllowed:  val.OriginAllowed,
		}
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   records,
	})
}
