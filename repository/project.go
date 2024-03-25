package repository

import (
	"yazmeyaa_projects/model"

	"gorm.io/gorm"
)

type ProjectsRepositoryImpl struct {
	Db *gorm.DB
}

type ProjectRepository interface {
	Save(projects model.Project)
	Update(projects model.Project)
	Delete(projectId int)
	FindById(projectId int) (project model.Project, err error)
	FindAll() []model.Project
}
