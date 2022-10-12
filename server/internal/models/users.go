package models

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/asxraj/url-shortener/internal/validator"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrDuplicateEmail = errors.New("duplicate email")
)

var AnonymousUser = &User{}

type User struct {
	ID        int      `json:"id"`
	Firstname string   `json:"first_name"`
	Lastname  string   `json:"last_name"`
	Email     string   `json:"email"`
	Password  password `json:"-"`
	Activated bool     `json:"activated"`
}

func (u *User) IsAnonymous() bool {
	return u == AnonymousUser
}

type password struct {
	plaintext *string
	hash      []byte
}

type UserModel struct {
	DB *sql.DB
}

func (p *password) Set(plainttextPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainttextPassword), 12)
	if err != nil {
		return err
	}

	p.plaintext = &plainttextPassword
	p.hash = hash

	return nil
}

func (p *password) Matches(plaintextPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.hash, []byte(plaintextPassword))
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

func ValidateEmail(v *validator.Validator, email string) {
	v.Check(email != "", "email", "cannot be empty")
	v.Check(validator.Matches(email, validator.EmailRX), "email", "must be a valid email address")
}

func ValidatePasswordPlaintext(v *validator.Validator, password password) {
	v.Check(*password.plaintext != "", "password", "cannot be empty")
	v.Check(len([]rune(*password.plaintext)) >= 8, "password", "must be at least 8 characters long")
	v.Check(len([]rune(*password.plaintext)) <= 72, "password", "must not be more than 72 charachters long")
}

func ValidateUser(v *validator.Validator, user *User) {
	v.Check(user.Firstname != "", "first_name", "cannot be empty")
	v.Check(user.Lastname != "", "last_name", "cannot be empty")

	ValidateEmail(v, user.Email)

	if user.Password.plaintext != nil {
		ValidatePasswordPlaintext(v, user.Password)
	}

	if user.Password.hash == nil {
		panic("missing password hash for user")
	}
}

func (m UserModel) Insert(user *User) error {
	query := `
	    INSERT INTO users (first_name, last_name, email, hashed_password) 
        VALUES ($1,$2,$3,$4)
		RETURNING id
	`

	args := []any{user.Firstname, user.Lastname, user.Email, user.Password.hash}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&user.ID)
	if err != nil {
		switch {
		case err.Error() == `ERROR: duplicate key value violates unique constraint "users_email_key" (SQLSTATE 23505)`:
			return ErrDuplicateEmail
		default:
			return err
		}
	}

	return nil
}

func (m UserModel) GetUserByEmail(user *User) error {
	query := `
        SELECT id, first_name, last_name, email, hashed_password
        FROM users 
        WHERE email = $1
    `

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, user.Email).Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email, &user.Password.hash)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrRecordNotFound
		default:
			return err
		}
	}

	return nil
}

func (m UserModel) Get(id int) (*User, error) {
	query := `
        SELECT id, first_name, last_name, email, hashed_password
        FROM users 
        WHERE id = $1
    `

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var user User

	err := m.DB.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email, &user.Password.hash)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (m UserModel) SaveURL(user *User, longUrl, shortUrl string, expires time.Time) error {
	// Look up automatic deletion when expiry is less than now()
	query := `
	    INSERT INTO urls (user_id, long_url, short_url, expires) 
        VALUES ($1,$2,$3,$4)
	`

	args := []any{user.ID, longUrl, shortUrl, expires}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&user.ID)
	if err != nil {
		switch {
		case err.Error() == `ERROR: duplicate key value violates unique constraint "users_email_key" (SQLSTATE 23505)`:
			return ErrDuplicateEmail
		default:
			return err
		}
	}

	return nil
}
