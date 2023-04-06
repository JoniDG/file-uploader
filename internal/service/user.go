package service

import (
	"encoding/csv"
	"file-uploader/internal/defines"
	"file-uploader/internal/domain"
	"file-uploader/internal/repository"
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
	status = defines.StatusProcessing
	err = r.repoStatus.Create(fileName, status)
	if err != nil {
		e := &domain.Error{
			NameFile: fileName,
			Err:      err,
			Msg:      "Error creating status",
		}
		log.Println(e.Error())
	}
	for _, row := range *rows {
		user, err := domain.RowFileToUser(row)
		if err != nil {
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
			status = defines.StatusError
			updateErr := r.repoStatus.Update(fileName, status)
			if updateErr != nil {
				log.Println(
					&domain.Error{
						NameFile: fileName,
						Err:      updateErr,
						Msg:      "Error updating status",
					})
			}
			e := &domain.Error{
				NameFile: fileName,
				Err:      err,
				Msg:      "Error creating user",
			}
			log.Println(e.Error())
		}
	}
	status = defines.StatusOk
	err = os.Remove(fileName)
	if err != nil {
		log.Println(
			&domain.Error{
				NameFile: fileName,
				Err:      err,
				Msg:      "Error deleting file",
			})
	}

	err = r.repoStatus.Update(fileName, status)
	if err != nil {
		log.Println(
			&domain.Error{
				NameFile: fileName,
				Err:      err,
				Msg:      "Error updating status",
			})
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
