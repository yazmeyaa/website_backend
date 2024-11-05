package controller

import (
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

	if err := ctx.ShouldBindJSON(&credentails); err != nil {
		helper.HandleHTTPError(ctx, http.StatusBadRequest, err)
		return
	}

	user, err := controller.authService.CheckAuth(credentails)
	if err != nil {
		webResponse := response.ErrorResponse{
			Error: err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		ctx.Abort()
		return
	}

	token, err := helper.CreateAccessToken(controller.appConfig.JWT.Secret, user, int(time.Hour.Nanoseconds())*168)
	if err != nil {
		helper.HandleHTTPError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.Header("X-Token", token)
	ctx.Status(http.StatusNoContent)
}

func (controller *AuthController) Register(ctx *gin.Context) {
	credentails := request.AuthCredentails{}

	err := ctx.ShouldBindJSON(&credentails)
	if err != nil {
		webResponse := response.ErrorResponse{
			Error: err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		ctx.Abort()
		return
	}

	user, createUserError := controller.authService.Create(credentails)
	if createUserError != nil {
		webResponse := response.ErrorResponse{
			Error: createUserError.Error(),
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		ctx.Abort()
		return
	}

	webResponse := response.CreateUserResponse{
		ID:       user.ID,
		Username: user.Username,
	}
	ctx.JSON(http.StatusNoContent, webResponse)
}

func (controller *AuthController) ValidateJWT(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}
