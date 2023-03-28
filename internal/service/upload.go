package service

import (
	"encoding/csv"
	"file-uploader/internal/domain"
	"file-uploader/internal/repository"
	"fmt"
	"log"
	"os"
	"strconv"
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
	rows := ReadFile(file)
	for _, row := range rows {
		id, err := strconv.Atoi(row[0])
		if err != nil {
			log.Println(err)
		}
		dataFile := domain.File{
			ID:       int64(id),
			Name:     row[1],
			LastName: row[2],
			Email:    row[3],
			Job:      row[4],
		}
		r.repo.UploadFile(&dataFile)
	}
}

func ReadFile(file *os.File) [][]string {
	reader := csv.NewReader(file)
	reader.Comma = ','
	rows, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return nil
	}
	return rows
}
