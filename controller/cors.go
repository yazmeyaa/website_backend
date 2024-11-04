package controller

import (
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

}

func (cc *CorsController) EnableOrigin(ctx *gin.Context) {

}

func (cc *CorsController) DisableOrigin(ctx *gin.Context) {

}

func (cc *CorsController) RemoveOrigin(ctx *gin.Context) {

}

func (cc *CorsController) GetAllOrigins(ctx *gin.Context) {

}
