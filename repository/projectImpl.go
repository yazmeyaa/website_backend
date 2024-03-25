package repository

import (
	"errors"
	"yazmeyaa_projects/helper"
	"yazmeyaa_projects/model"

	"gorm.io/gorm"
)

func NewProjectsRepositoryImpl(Db *gorm.DB) *ProjectsRepositoryImpl {
	return &ProjectsRepositoryImpl{Db: Db}
}

func (p ProjectsRepositoryImpl) Save(project model.Project) {
	result := p.Db.Create(&project)
	helper.ErrorPanic(result.Error)
}

func (p ProjectsRepositoryImpl) Update(project model.Project) {
	result := p.Db.Model(&project).Updates(project)
	helper.ErrorPanic(result.Error)
}

func (p ProjectsRepositoryImpl) Delete(projectId int) {
	var project model.Project

	result := p.Db.Where("ID = ?", projectId).Delete(&project)
	helper.ErrorPanic(result.Error)
}

func (p ProjectsRepositoryImpl) FindById(projectId int) (model.Project, error) {
	var project model.Project

	result := p.Db.Find(&project, projectId)

	if result != nil {
		return project, nil
	} else {
		return project, errors.New("tag is not found")
	}
}

func (p ProjectsRepositoryImpl) FindAll() []model.Project {
	var tags []model.Project
	result := p.Db.Find(&tags)
	helper.ErrorPanic(result.Error)
	return tags
}
