package models

type User struct {
	Username string `db:"username"`
	Email    string `db:"email"`
	Name     string `db:"name"`
	Password string `db:"password"`
}
