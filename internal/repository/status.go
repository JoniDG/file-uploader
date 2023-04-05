package repository

import (
	"file-uploader/internal/defines"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type StatusRepository interface {
	CreateStatus(*string, *string) error
	UpdateStatus(*string, *string) error
}

type statusRepository struct {
	db         *sqlx.DB
	sqlBuilder *tableSQL2
}

func NewStatusFileRepo(db *sqlx.DB) StatusRepository {
	return &statusRepository{
		db: db,
		sqlBuilder: &tableSQL2{
			table: defines.TableStatusFile,
		},
	}
}

func (r *statusRepository) CreateStatus(nameFile *string, status *string) error {
	query, args, err := r.sqlBuilder.CreateSQL(nameFile, status)
	if err != nil {
		return err
	}
	_, err = r.db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r *statusRepository) UpdateStatus(nameFile *string, status *string) error {
	query, args, err := r.sqlBuilder.UpdateSQL(nameFile, status)
	if err != nil {
		return err
	}
	_, err = r.db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

type tableSQL2 struct {
	table string
}

func (s *tableSQL2) CreateSQL(nameFile *string, status *string) (string, []interface{}, error) {
	query, args, err := squirrel.Insert(s.table).
		Columns("fileName", "status").
		Values(nameFile, status).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	return query, args, err
}

func (s *tableSQL2) UpdateSQL(nameFile *string, status *string) (string, []interface{}, error) {
	query, args, err := squirrel.Update(s.table).
		Set("status", status).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{"fileName": nameFile}).
		ToSql()
	return query, args, err
}
