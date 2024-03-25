package service

import (
	"yazmeyaa_projects/data/request"
	"yazmeyaa_projects/data/response"
	"yazmeyaa_projects/helper"
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

func (service *ProjectsServiceImpl) Create(project request.CreateProjectRequest) {
	err := service.Validate.Struct(project)
	helper.ErrorPanic(err)
	tagModel := model.Project{
		Name:        project.Name,
		Href:        &project.Href,
		Description: project.Description,
		Img:         project.Img,
		GithubUrl:   &project.GithubUrl,
	}
	service.ProjectsRepository.Save(tagModel)
}

func (service *ProjectsServiceImpl) Update(project request.UpdateProjectRequest) {
	err := service.Validate.Struct(project)
	helper.ErrorPanic(err)
	tagModel := model.Project{
		Name:        project.Name,
		Href:        &project.Href,
		Description: project.Description,
		Img:         project.Img,
		GithubUrl:   &project.GithubUrl,
	}
	service.ProjectsRepository.Update(tagModel)
}

func (service *ProjectsServiceImpl) Delete(projectId int) {
	service.ProjectsRepository.Delete(projectId)
}

func (service *ProjectsServiceImpl) FindById(projectId int) response.ProjectsResponse {
	result, err := service.ProjectsRepository.FindById(projectId)
	if err != nil {
		helper.ErrorPanic(err)
	}

	resp := response.ProjectsResponse{
		Name:        result.Name,
		Href:        result.Href,
		Description: result.Description,
		Img:         result.Img,
		GithubUrl:   result.GithubUrl,
	}

	return resp
}

func (service *ProjectsServiceImpl) FindAll() []response.ProjectsResponse {
	result := service.ProjectsRepository.FindAll()

	var projects []response.ProjectsResponse
	for _, value := range result {
		project := response.ProjectsResponse{
			ID:          value.ID,
			Name:        value.Name,
			Href:        value.Href,
			Description: value.Description,
			Img:         value.Img,
			GithubUrl:   value.GithubUrl,
		}
		projects = append(projects, project)
	}

	return projects
}
