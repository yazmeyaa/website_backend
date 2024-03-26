package repository

import (
	"yazmeyaa_projects/model"
)

type ProjectRepository interface {
	Save(projects model.Project)
	Update(projects model.Project)
	Delete(projectId int)
	FindById(projectId int) (project model.Project, err error)
	FindAll() []model.Project
}
