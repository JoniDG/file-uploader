package service

import (
	"encoding/csv"
	"file-uploader/internal/repository"
	"fmt"
	"log"
	"os"
)

type UploadService interface {
	UploadFile(string)
}
type uploadService struct {
	repo repository.UploadRepository
}

func NewUploadService(repo repository.UploadRepository) UploadService {
	return &uploadService{repo: repo}
}

func (r *uploadService) UploadFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Cannot open '%s': %s\n", fileName, err.Error())
	}
	defer file.Close()
	ReadFile(file)
}

func ReadFile(file *os.File) {
	r := csv.NewReader(file)
	rows, err := r.ReadAll()
	if err != nil {
		fmt.Println("Error al leer las filas del archivo:", err)
		return
	}
	log.Println(rows[0][1])
}
