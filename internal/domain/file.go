package domain

type File struct {
	ID       int64  `db:"id"`
	Name     string `db:"name"`
	LastName string `db:"last_name"`
	Email    string `db:"email"`
	Job      string `db:"job"`
}
