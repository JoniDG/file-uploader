package controller

import (
	"file-uploader/internal/service"
	"log"
)

type UploadController interface {
	UploadFile(string)
}
type uploadController struct {
	svc service.UploadService
}

func NewUploadController(svc service.UploadService) UploadController {
	return &uploadController{svc: svc}
}

func (c *uploadController) UploadFile(fileName string) {
	log.Printf("Nombre del archivo: %+v\n", fileName)
}
