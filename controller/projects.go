package controller

import (
	"errors"
	"net/http"
	"strconv"
	"yazmeyaa_projects/data/request"
	"yazmeyaa_projects/data/response"
	"yazmeyaa_projects/helper"
	"yazmeyaa_projects/repository"
	"yazmeyaa_projects/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ProjectsController struct {
	projectsService service.ProjectsService
	validate        *validator.Validate
}

func NewProjectsController(service service.ProjectsService, validate *validator.Validate) *ProjectsController {
	return &ProjectsController{
		projectsService: service,
		validate:        validate,
	}
}

func (controller *ProjectsController) Update(ctx *gin.Context) {
	updateProjectsRequest := request.UpdateProjectRequest{}
	if err := ctx.ShouldBindJSON(&updateProjectsRequest); err != nil {
		helper.HandleHTTPError(ctx, http.StatusBadRequest, err)
		return
	}

	project, err := controller.projectsService.FindById(updateProjectsRequest.ID)
	if err != nil {
		helper.HandleHTTPError(ctx, http.StatusBadRequest, err)
		return
	}

	project.Description = updateProjectsRequest.Description
	project.GithubUrl = &updateProjectsRequest.GithubUrl
	project.Href = &updateProjectsRequest.Href
	project.Img = updateProjectsRequest.Img
	project.ImgUrl = &updateProjectsRequest.ImgUrl
	project.Name = updateProjectsRequest.Name

	if err := controller.projectsService.Update(project); err != nil {
		helper.HandleHTTPError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.Header("Content-Type", "application/json")
	ctx.Status(http.StatusNoContent)
}

func (controller *ProjectsController) Create(ctx *gin.Context) {
	createProjectRequest := request.CreateProjectRequest{}
	if err := ctx.ShouldBindJSON(&createProjectRequest); err != nil {
		helper.HandleHTTPError(ctx, http.StatusBadRequest, err)
		return
	}

	project, err := controller.projectsService.Create(repository.CreateProjectData{
		Name:        createProjectRequest.Name,
		Description: createProjectRequest.Description,
		Href:        createProjectRequest.Href,
		Img:         createProjectRequest.Img,
		GithubUrl:   createProjectRequest.GithubUrl,
		ImgUrl:      createProjectRequest.ImgUrl,
	})

	if err != nil {
		helper.HandleHTTPError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, project.ToHTTPResponse())
}

func (controller *ProjectsController) Delete(ctx *gin.Context) {
	projectId := ctx.Param("id")
	id, err := strconv.Atoi(projectId)
	if err != nil {
		helper.HandleHTTPError(ctx, http.StatusBadRequest, errors.New("unexpected value in \"id\" param"))
		return
	}

	if err := controller.projectsService.Delete(id); err != nil {
		helper.HandleHTTPError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (controller *ProjectsController) GetAll(ctx *gin.Context) {
	projects, err := controller.projectsService.FindAll()
	if err != nil {
		helper.HandleHTTPError(ctx, http.StatusInternalServerError, err)
		return
	}

	resp := make([]response.ProjectsResponse, len(projects))

	for idx, value := range projects {
		resp[idx] = value.ToHTTPResponse()
	}

	ctx.JSON(http.StatusOK, resp)
}

func (controller *ProjectsController) GetById(ctx *gin.Context) {
	projectId := ctx.Param("id")
	id, err := strconv.Atoi(projectId)

	if err != nil {
		helper.HandleHTTPError(ctx, http.StatusBadRequest, err)
		return
	}

	project, err := controller.projectsService.FindById(id)
	if err != nil {
		helper.HandleHTTPError(ctx, http.StatusBadRequest, err)
		return
	}
	response := project.ToHTTPResponse()

	ctx.JSON(http.StatusOK, response)
}
