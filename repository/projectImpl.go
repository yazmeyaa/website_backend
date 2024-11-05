package repository

import (
	"yazmeyaa_projects/model"

	"gorm.io/gorm"
)

type ProjectsRepositoryImpl struct {
	Db *gorm.DB
}

func NewProjectsRepositoryImpl(Db *gorm.DB) ProjectRepository {
	return &ProjectsRepositoryImpl{Db: Db}
}

func (p ProjectsRepositoryImpl) Save(data CreateProjectData) (*model.Project, error) {
	var project model.Project

	project.Name = data.Name
	project.Description = data.Description
	project.Href = data.Href
	project.GithubUrl = data.GithubUrl
	project.Img = data.Img
	project.ImgUrl = data.ImgUrl

	if err := p.Db.Create(&project).Error; err != nil {
		return nil, err
	}

	return &project, nil
}

func (p ProjectsRepositoryImpl) Update(project *model.Project) error {
	return p.Db.Model(project).Where("ID = ?", project.ID).Updates(project).Error
}

func (p ProjectsRepositoryImpl) Delete(projectId int) error {
	var project model.Project

	return p.Db.Where("ID = ?", projectId).Delete(&project).Error
}

func (p ProjectsRepositoryImpl) FindById(projectId int) (*model.Project, error) {
	var project model.Project

	if err := p.Db.Find(&project, projectId).Error; err != nil {
		return nil, err
	}

	return &project, nil
}

func (p ProjectsRepositoryImpl) FindAll() ([]model.Project, error) {
	var tags []model.Project
	if err := p.Db.Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}
