package service

import (
	"encoding/csv"
	"file-uploader/internal/domain"
	"file-uploader/internal/repository"
	"fmt"
	"log"
	"os"
)

type FileService interface {
	HandlerFile(string)
}
type fileService struct {
	repo repository.UserRepository
}

func NewFileService(repo repository.UserRepository) FileService {
	return &fileService{repo: repo}
}

func (r *fileService) HandlerFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Cannot open '%s': %s\n", fileName, err.Error())
	}
	defer file.Close()
	rows := ReadFile(file)
	for _, row := range rows {
		user, err := domain.RowFileToUser(row)
		if err != nil {
			fmt.Printf("Error ID format: %+v\n", row)
			continue
		}
		err = r.repo.Create(user)
		if err != nil {
			log.Println(err)
		}
	}
	fmt.Printf("Archivo %s cargado\n", fileName)
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
