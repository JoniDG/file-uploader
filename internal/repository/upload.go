package repository

type UploadRepository interface {
	UploadFile()
}
type uploadRepository struct {
}

func NewUploadRepository() UploadRepository {
	return &uploadRepository{}
}

func (r *uploadRepository) UploadFile() {

}
