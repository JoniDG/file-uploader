package repository

import (
	"file-uploader/internal/defines"
	"file-uploader/internal/domain"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"log"
)

type UploadRepository interface {
	UploadFile(file *domain.File)
}
type uploadRepository struct {
	db         *sqlx.DB
	sqlBuilder tableSQL
}

func NewUploadRepository(db *sqlx.DB) UploadRepository {
	return &uploadRepository{
		db: db,
		sqlBuilder: tableSQL{
			table: defines.TableUploadFile,
		},
	}
}

func (r *uploadRepository) UploadFile(row *domain.File) {
	query, args, err := r.sqlBuilder.CreateSQL(row)
	if err != nil {
		log.Println(err)
	}
	_, err = r.db.Exec(query, args...)
	if err != nil {
		log.Println(err)
	}
}

type tableSQL struct {
	table string
}

func (s *tableSQL) CreateSQL(p *domain.File) (string, []interface{}, error) {
	query, args, err := squirrel.Insert(s.table).
		Columns("id", "name", "lastname", "email", "job").
		Values(p.ID, p.Name, p.LastName, p.Email, p.Job).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	return query, args, err
}
