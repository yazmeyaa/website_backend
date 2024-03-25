package service

import (
	"yazmeyaa_projects/data/request"
	"yazmeyaa_projects/data/response"
)

type ProjectsService interface {
	Create(project request.CreateProjectRequest)
	Update(project request.UpdateProjectRequest)
	Delete(projectId int)
	FindById(projectId int) response.ProjectsResponse
	FindAll() []response.ProjectsResponse
}
