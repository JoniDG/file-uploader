package service

import "file-uploader/internal/repository"

type UploadService interface {
	UploadFile()
}
type uploadService struct {
	repo repository.UploadRepository
}

func NewUploadService(repo repository.UploadRepository) UploadService {
	return &uploadService{repo: repo}
}

func (r *uploadService) UploadFile() {
}
