package controller

import (
	"fmt"
	"net/http"
	"time"
	"yazmeyaa_projects/config"
	"yazmeyaa_projects/data/request"
	"yazmeyaa_projects/data/response"
	"yazmeyaa_projects/helper"
	"yazmeyaa_projects/service"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService service.AuthService
	appConfig   config.AppConfig
}

func NewAuthController(service service.AuthService) *AuthController {
	return &AuthController{
		authService: service,
		appConfig:   *config.NewAppConfig(),
	}
}

func (controller *AuthController) Login(ctx *gin.Context) {
	credentails := request.AuthCredentails{}

	err := ctx.ShouldBindJSON(&credentails)
	helper.ErrorPanic(err)

	user, err := controller.authService.CheckAuth(credentails)
	fmt.Println(user.Username)
	if err != nil {
		webResponse := response.Response{
			Code:   http.StatusBadRequest,
			Data:   err.Error(),
			Status: "Bad Request",
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		ctx.Abort()
		return
	}

	token, err := helper.CreateAccessToken(controller.appConfig.JWT.Secret, int(time.Hour.Nanoseconds())*168)
	helper.ErrorPanic(err)

	ctx.Header("X-Token", token)
}

func (controller *AuthController) Register(ctx *gin.Context) {
	credentails := request.AuthCredentails{}

	err := ctx.ShouldBindJSON(&credentails)
	if err != nil {
		webResponse := response.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad request",
			Data:   nil,
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		ctx.Abort()
		return
	}

	user, createUserError := controller.authService.Create(credentails)
	if createUserError != nil {
		webResponse := response.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad request",
			Data:   createUserError.Error(),
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		ctx.Abort()
		return
	}

	webResponse := response.Response{
		Code:   http.StatusNoContent,
		Status: "Created",
		Data:   user,
	}
	ctx.JSON(http.StatusNoContent, webResponse)
}
