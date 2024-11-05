package repository

import (
	"yazmeyaa_projects/model"
)

type CreateProjectData struct {
	Name        string
	Description string
	Href        *string
	Img         string
	GithubUrl   *string
	ImgUrl      *string
}

type ProjectRepository interface {
	Save(projects CreateProjectData) (*model.Project, error)
	Update(projects *model.Project) error
	Delete(projectId int) error
	FindById(projectId int) (project *model.Project, err error)
	FindAll() ([]model.Project, error)
}
