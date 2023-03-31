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
	HandlerFile(string) error
}
type fileService struct {
	repo repository.UserRepository
}

func NewFileService(repo repository.UserRepository) FileService {
	return &fileService{repo: repo}
}

func (r *fileService) HandlerFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return &domain.Error{
			NameFile: fileName,
			Err:      err,
			Msg:      "Cannot open file",
		}
	}
	defer file.Close()
	rows, err := ReadFile(file)
	if err != nil {
		return &domain.Error{
			NameFile: fileName,
			Err:      err,
			Msg:      "Error al leer el archivo",
		}
	}
	for _, row := range rows {
		user, err := domain.RowFileToUser(row)
		if err != nil {
			e := domain.Error{
				NameFile: fileName,
				Err:      err,
				Msg:      fmt.Sprintf("Error ID format: %+v\n", row),
			}
			log.Println(e.Error())
			continue
		}
		err = r.repo.Create(user)
		if err != nil {
			e := domain.Error{
				NameFile: fileName,
				Err:      err,
				Msg:      "Error creating user",
			}
			log.Println(e.Error())
		}
	}
	return nil
}

func ReadFile(file *os.File) ([][]string, error) {
	reader := csv.NewReader(file)
	reader.Comma = ','
	rows, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}
	return rows, nil
}
