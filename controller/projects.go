package controller

import (
	"net/http"
	"strconv"
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

func (controller *ProjectsController) Update(ctx *gin.Context) {
	updateProjectsRequest := request.UpdateProjectRequest{}
	err := ctx.ShouldBindJSON(&updateProjectsRequest)
	helper.ErrorPanic(err)

	tagId := ctx.Param("id")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	updateProjectsRequest.ID = id

	controller.projectsService.Update(updateProjectsRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
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

func (controller *ProjectsController) Delete(ctx *gin.Context) {
	projectId := ctx.Param("id")
	id, err := strconv.Atoi(projectId)
	helper.ErrorPanic(err)

	controller.projectsService.Delete(id)

	webResponse := response.Response{
		Code:   http.StatusNoContent,
		Status: "Ok",
		Data:   nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusNoContent, webResponse)
}

func (controller *ProjectsController) GetAll(ctx *gin.Context) {
	projectsResponse := controller.projectsService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   projectsResponse,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ProjectsController) GetById(ctx *gin.Context) {
	projectId := ctx.Param("id")
	id, err := strconv.Atoi(projectId)
	helper.ErrorPanic(err)

	project := controller.projectsService.FindById(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   project,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
