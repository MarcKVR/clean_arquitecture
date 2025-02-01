package service

import (
	"mime/multipart"

	"github.com/MarcKVR/clean_arquitecture/domain/repositories"
)

type FileService struct {
	repo repositories.FileRepository
}

func NewFileService(repo repositories.FileRepository) *FileService {
	return &FileService{
		repo: repo,
	}
}

func (s *FileService) UploadFile(file *multipart.FileHeader) (string, error) {
	destination := "uploads/" + file.Filename
	return s.repo.UploadFile(file, destination)
}
