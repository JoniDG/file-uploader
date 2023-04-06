package repository

import (
	"file-uploader/internal/defines"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type StatusRepository interface {
	Create(nameFile string, status string, createdAt string) error
	Update(nameFile string, status string, uodatedAt string) error
}

type statusRepository struct {
	db         *sqlx.DB
	sqlBuilder *tableStatus
}

func NewStatusFileRepo(db *sqlx.DB) StatusRepository {
	return &statusRepository{
		db: db,
		sqlBuilder: &tableStatus{
			table: defines.TableStatusFile,
		},
	}
}

func (r *statusRepository) Create(nameFile string, status string, createdAt string) error {
	query, args, err := r.sqlBuilder.CreateSQL(nameFile, status, createdAt)
	if err != nil {
		return err
	}
	_, err = r.db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r *statusRepository) Update(nameFile string, status string, updatedAt string) error {
	query, args, err := r.sqlBuilder.UpdateSQL(nameFile, status, updatedAt)
	if err != nil {
		return err
	}
	_, err = r.db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

type tableStatus struct {
	table string
}

func (s *tableStatus) CreateSQL(nameFile string, status string, createdAt string) (string, []interface{}, error) {
	query, args, err := squirrel.Insert(s.table).
		Columns("fileName", "status", "created_at").
		Values(nameFile, status, createdAt).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	return query, args, err
}

func (s *tableStatus) UpdateSQL(nameFile string, status string, updatedAt string) (string, []interface{}, error) {
	query, args, err := squirrel.Update(s.table).
		Set("status", status).
		Set("updated_at", updatedAt).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{"fileName": nameFile}).
		ToSql()
	return query, args, err
}
