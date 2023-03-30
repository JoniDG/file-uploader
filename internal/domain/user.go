package domain

import (
	"strconv"
)

type User struct {
	ID       uint64 `db:"id"`
	Name     string `db:"name"`
	LastName string `db:"last_name"`
	Email    string `db:"email"`
	Job      string `db:"job"`
}

func RowFileToUser(row []string) (*User, error) {
	id, err := strconv.ParseUint(row[0], 10, 64)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:       id,
		Name:     row[1],
		LastName: row[2],
		Email:    row[3],
		Job:      row[4],
	}, nil
}
