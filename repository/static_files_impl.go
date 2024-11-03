package repository

import (
	"yazmeyaa_projects/model"

	"gorm.io/gorm"
)

type staticFilesRepository struct {
	db *gorm.DB
}

// GetAll implements StaticFileRepository.
func (s *staticFilesRepository) GetAll() ([]model.StaticFile, error) {
	var files []model.StaticFile
	result := s.db.Find(&files)

	return files, result.Error
}

// Create implements StaticFileRepository.
func (s *staticFilesRepository) Create(fileName string) (*model.StaticFile, error) {
	file := model.StaticFile{
		FileName: fileName,
	}
	if err := s.db.Create(&file).Error; err != nil {
		return nil, err
	}
	return &file, nil
}

// Delete implements StaticFileRepository.
func (s *staticFilesRepository) Delete(id int) error {
	return s.db.Where("id = ?", id).Delete(&model.StaticFile{}).Error
}

// GetById implements StaticFileRepository.
func (s *staticFilesRepository) GetById(id int) (*model.StaticFile, error) {
	file := model.StaticFile{}

	if err := s.db.Where("id = ?", id).First(&file).Error; err != nil {
		return nil, err
	}

	return &file, nil
}

func (s *staticFilesRepository) GetByFileName(filename string) (*model.StaticFile, error) {
	file := model.StaticFile{}

	if err := s.db.Where("file_name = ?", filename).First(&file).Error; err != nil {
		return nil, err
	}

	return &file, nil
}

func NewStaticFileRepository(db *gorm.DB) StaticFileRepository {
	return &staticFilesRepository{
		db: db,
	}
}
