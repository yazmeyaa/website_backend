package service

import (
	"yazmeyaa_projects/model"
	"yazmeyaa_projects/repository"
)

type ProjectsService interface {
	Create(project repository.CreateProjectData) (*model.Project, error)
	Update(project *model.Project) error
	Delete(projectId int) error
	FindById(projectId int) (*model.Project, error)
	FindAll() ([]model.Project, error)
}
