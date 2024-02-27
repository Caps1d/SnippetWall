package models

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

type UserModel struct {
	DB *pgxpool.Pool
}

func (m *UserModel) Insert(name, email, password string) error {
	// encrypt user password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	query := `
            INSERT INTO users (name, email, hashed_password, created)
            VALUES ($1, $2, $3, CURRENT_TIMESTAMP)`

	_, err = m.DB.Exec(context.Background(), query, name, email, string(hash))
	if err != nil {
		var pgError *pgconn.PgError
		if errors.As(err, &pgError) {
			code, _ := strconv.Atoi(pgError.Code)
			if code == 23505 && strings.Contains(pgError.Message, "users_uc_email") {
				return ErrDuplicateEmail
			}
		}
	}

	return err
}

func (m *UserModel) Authenticate(name, email, password string) (int, error) {
	return 0, nil
}
func (m *UserModel) Exists(id int) (bool, error) {
	return false, nil
}
