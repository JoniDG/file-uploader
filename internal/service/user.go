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
	repoUser   repository.UserRepository
	repoStatus repository.StatusRepository
}

func NewFileService(repoUser repository.UserRepository, repoStatus repository.StatusRepository) FileService {
	return &fileService{
		repoUser:   repoUser,
		repoStatus: repoStatus,
	}
}

func (r *fileService) HandlerFile(fileName string) error {
	file, err := os.Open(fileName)
	var status string
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
			Msg:      "Error reading file",
		}
	}
	status = "Processing"
	err = r.repoStatus.CreateStatus(&fileName, &status)
	if err != nil {
		fmt.Printf("Error en status: %v\n", err)
	}
	for _, row := range *rows {
		user, err := domain.RowFileToUser(row)
		if err != nil {
			status = "Processing with error"
			failStatus := r.repoStatus.UpdateStatus(&fileName, &status)
			if failStatus != nil {
				fmt.Printf("Error en status: %v\n", err)
			}
			e := &domain.Error{
				NameFile: fileName,
				Err:      err,
				Msg:      "Error parsing row to user",
			}
			log.Println(e.Error())
			continue
		}
		err = r.repoUser.Create(user)
		if err != nil {
			status = "Error"
			failStatus := r.repoStatus.UpdateStatus(&fileName, &status)
			if failStatus != nil {
				fmt.Printf("Error en status: %v\n", err)
			}
			e := &domain.Error{
				NameFile: fileName,
				Err:      err,
				Msg:      "Error creating user",
			}
			log.Println(e.Error())
		}
	}
	if status == "Processing with error" {
		status = "Processed with error"
	} else {
		status = "Processed"
		err = os.Remove(fileName)
		if err != nil {
			return &domain.Error{
				NameFile: fileName,
				Err:      err,
				Msg:      "Cannot delete file",
			}
		}
	}
	err = r.repoStatus.UpdateStatus(&fileName, &status)
	if err != nil {
		fmt.Printf("Error en status: %v\n", err)
	}
	return nil
}

func ReadFile(file *os.File) (*[][]string, error) {
	reader := csv.NewReader(file)
	reader.Comma = ','
	rows, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}
	return &rows, nil
}
