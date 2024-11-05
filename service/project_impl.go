package service

import (
	"yazmeyaa_projects/model"
	"yazmeyaa_projects/repository"

	"github.com/go-playground/validator/v10"
)

type ProjectsServiceImpl struct {
	ProjectsRepository repository.ProjectRepository
	Validate           *validator.Validate
}

func NewProjectsServiceImpl(projectsRepository repository.ProjectRepository, validate *validator.Validate) ProjectsService {
	return &ProjectsServiceImpl{
		Validate:           validate,
		ProjectsRepository: projectsRepository,
	}
}

func (service *ProjectsServiceImpl) Create(data repository.CreateProjectData) (*model.Project, error) {
	return service.ProjectsRepository.Save(data)
}

func (service *ProjectsServiceImpl) Update(project *model.Project) error {
	return service.ProjectsRepository.Update(project)
}

func (service *ProjectsServiceImpl) Delete(projectId int) error {
	return service.ProjectsRepository.Delete(projectId)
}

func (service *ProjectsServiceImpl) FindById(projectId int) (*model.Project, error) {
	result, err := service.ProjectsRepository.FindById(projectId)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *ProjectsServiceImpl) FindAll() ([]model.Project, error) {
	return service.ProjectsRepository.FindAll()
}
