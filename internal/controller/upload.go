package controller

import (
	"file-uploader/internal/service"
	"log"
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
	err := c.svc.HandlerFile(*fileName)
	if err != nil {
		log.Println(err.Error())
	}
}
