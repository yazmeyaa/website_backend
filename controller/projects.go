package controller

import (
	"net/http"
	"strconv"
	"yazmeyaa_projects/data/request"
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

	controller.projectsService.Update(updateProjectsRequest)

	ctx.Header("Content-Type", "application/json")
	ctx.Status(http.StatusNoContent)
}

func (controller *ProjectsController) Create(ctx *gin.Context) {
	createProjectRequest := request.CreateProjectRequest{}
	err := ctx.ShouldBindJSON(&createProjectRequest)

	helper.ErrorPanic(err)

	controller.projectsService.Create(createProjectRequest)

	ctx.Header("Content-Type", "application/json")
	ctx.Status(http.StatusNoContent)
}

func (controller *ProjectsController) Delete(ctx *gin.Context) {
	projectId := ctx.Param("id")
	id, err := strconv.Atoi(projectId)
	helper.ErrorPanic(err)

	controller.projectsService.Delete(id)

	ctx.Header("Content-Type", "application/json")
	ctx.Status(http.StatusNoContent)
}

func (controller *ProjectsController) GetAll(ctx *gin.Context) {
	projectsResponse := controller.projectsService.FindAll()

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, projectsResponse)
}

func (controller *ProjectsController) GetById(ctx *gin.Context) {
	projectId := ctx.Param("id")
	id, err := strconv.Atoi(projectId)
	helper.ErrorPanic(err)

	project := controller.projectsService.FindById(id)

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, project)
}
