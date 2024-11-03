package service

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
	"yazmeyaa_projects/model"
	"yazmeyaa_projects/repository"
)

const uploadDir = "uploads"

type staticFileService struct {
	repo repository.StaticFileRepository
}

func (s *staticFileService) Create(filename string, data []byte) (*model.StaticFile, error) {
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("could not create upload directory: %w", err)
	}

	uniqueFilename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), filename)
	filePath := filepath.Join(uploadDir, uniqueFilename)

	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return nil, fmt.Errorf("could not save file: %w", err)
	}

	file, err := s.repo.Create(uniqueFilename)
	if err != nil {
		return nil, fmt.Errorf("could not save file record in database: %w", err)
	}

	return file, nil
}

func (s *staticFileService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *staticFileService) GetById(id int) (*model.StaticFile, error) {
	return s.repo.GetById(id)
}

func (s *staticFileService) GetByFileName(fileName string) (*model.StaticFile, error) {
	return s.repo.GetByFileName(fileName)
}

func NewStaticFileService(repo repository.StaticFileRepository) StaticFileService {
	return &staticFileService{
		repo: repo,
	}
}
