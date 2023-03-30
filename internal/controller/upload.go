package controller

import (
	"file-uploader/internal/service"
)

type UploadController interface {
	UploadFile(*string)
}
type uploadController struct {
	svc service.FileService
}

func NewUploadController(svc service.FileService) UploadController {
	return &uploadController{svc: svc}
}

func (c *uploadController) UploadFile(fileName *string) {
	c.svc.HandlerFile(*fileName)
}
