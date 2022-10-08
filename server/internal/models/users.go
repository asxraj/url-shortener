package models

import (
	"database/sql"
	"time"
)

type User struct {
	Firstname string    `json:"first_name"`
	Lastname  string    `json:"last_name"`
	Username  string    `json:"user_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type UserModel struct {
	DB *sql.DB
}

func (m UserModel) Insert(user *User) error {
	// query := `
	//     INSERT INTO users (first_name, last_name, username, email, password_hash) VALUES ($1,$2,$3,$4,$5)

	// `

	return nil

}
