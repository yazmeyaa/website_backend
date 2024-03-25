package controller

import (
	"net/http"
	"yazmeyaa_projects/data/request"
	"yazmeyaa_projects/data/response"
	"yazmeyaa_projects/helper"
	"yazmeyaa_projects/service"

	"github.com/gin-gonic/gin"
)

type ProjectsController struct {
	projectsService service.ProjectsService
}

func NewProjectsController(service service.ProjectsService) *ProjectsController {
	return &ProjectsController{
		projectsService: service,
	}
}

func (controller *ProjectsController) Create(ctx *gin.Context) {
	createProjectRequest := request.CreateProjectRequest{}
	err := ctx.ShouldBindJSON(&createProjectRequest)

	helper.ErrorPanic(err)

	controller.projectsService.Create(createProjectRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
