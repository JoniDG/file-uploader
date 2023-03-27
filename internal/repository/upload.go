package repository

import (
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
	sqlBuilder personasSQL
}

func NewUploadRepository(db *sqlx.DB) UploadRepository {
	return &uploadRepository{
		db: db,
		sqlBuilder: personasSQL{
			table: "file.personas",
		},
	}
}

func (r *uploadRepository) UploadFile(row *domain.File) {
	query, args, err := r.sqlBuilder.CreateSQL(row)
	if err != nil {
		log.Println(err)
	}
	_, err = r.db.Exec(query, args...)
}

type personasSQL struct {
	table string
}

func (s *personasSQL) CreateSQL(p *domain.File) (string, []interface{}, error) {
	query, args, err := squirrel.Insert(s.table).
		Columns("id", "name", "lastname", "email", "job").
		Values(p.ID, p.Name, p.LastName, p.Email, p.Job).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	return query, args, err
}
