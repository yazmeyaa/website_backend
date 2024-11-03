package service

import "yazmeyaa_projects/model"

type StaticFileService interface {
	Create(filename string, data []byte) (*model.StaticFile, error)
	GetByFileName(fileName string) (*model.StaticFile, error)
	GetById(id int) (*model.StaticFile, error)
	Delete(id int) error
}
