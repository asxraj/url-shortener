package models

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/asxraj/url-shortener/internal/validator"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Firstname      string    `json:"first_name"`
	Lastname       string    `json:"last_name"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	hashedPassword []byte    `json:"-"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
}

type UserModel struct {
	DB *sql.DB
}

func (u *User) Matches(plaintextPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(u.hashedPassword, []byte(plaintextPassword))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}
	return true, nil
}

func (u *User) Set() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	if err != nil {
		return err
	}
	u.hashedPassword = hash
	return nil
}

func ValidateEmail(v *validator.Validator, email string) {
	v.Check(email != "", "email", "cannot be empty")
	v.Check(validator.Matches(email, validator.EmailRX), "email", "must be a valid email address")
}

func ValidatePassword(v *validator.Validator, password string) {
	v.Check(password != "", "password", "cannot be empty")
	v.Check(len([]rune(password)) >= 8, "password", "must be at least 8 characters long")
	v.Check(len([]rune(password)) <= 72, "password", "must not be more than 72 charachters long")
}

func ValidateUser(v *validator.Validator, user *User) {
	v.Check(user.Firstname != "", "first_name", "cannot be empty")
	v.Check(user.Lastname != "", "last_name", "cannot be empty")
	v.Check(user.Username != "", "username", "cannot be empty")
	v.Check(len([]rune(user.Username)) <= 16, "name", "must not be more than 16 characters long")

	ValidateEmail(v, user.Email)
	ValidatePassword(v, user.Password)
}

func (m UserModel) Insert(user *User) error {
	query := `
	    INSERT INTO users (first_name, last_name, username, email, password_hash) 
        VALUES ($1,$2,$3,$4,$5)
	`

	args := []any{user.Firstname, user.Lastname, user.Username, user.Email, user.hashedPassword}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
