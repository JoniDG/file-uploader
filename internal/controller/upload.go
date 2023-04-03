package controller

import (
	"file-uploader/internal/service"
)

type UploadController interface {
	UploadFile(*string) error
}
type uploadController struct {
	svc service.FileService
}

func NewUploadController(svc service.FileService) UploadController {
	return &uploadController{svc: svc}
}

func (c *uploadController) UploadFile(fileName *string) error {
	err := c.svc.HandlerFile(*fileName)
	if err != nil {
		return err
	}
	return nil
}
