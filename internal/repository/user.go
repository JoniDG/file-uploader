package repository

import (
	"file-uploader/internal/defines"
	"file-uploader/internal/domain"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	Create(file *domain.User) error
}
type userRepository struct {
	db         *sqlx.DB
	sqlBuilder *tableSQL
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{
		db: db,
		sqlBuilder: &tableSQL{
			table: defines.TableUploadFile,
		},
	}
}

func (r *userRepository) Create(user *domain.User) error {
	query, args, err := r.sqlBuilder.CreateSQL(user)
	if err != nil {
		return err
	}
	_, err = r.db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

type tableSQL struct {
	table string
}

func (s *tableSQL) CreateSQL(p *domain.User) (string, []interface{}, error) {
	query, args, err := squirrel.Insert(s.table).
		Columns("id", "name", "lastname", "email", "job").
		Values(p.ID, p.Name, p.LastName, p.Email, p.Job).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	return query, args, err
}
