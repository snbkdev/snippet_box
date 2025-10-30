package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID int
	Name string
	Email string
	HashedPassword []byte
	Created time.Time
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name, email, password string) error {
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `insert into users (name, email, hashed_password, created) values($1, $2, $3, now())`

	_, err = m.DB.Exec(stmt, name, email, hashed_password)
	if err != nil {
    	var pgErr *pq.Error
    	if errors.As(err, &pgErr) {
        	if pgErr.Code.Name() == "unique_violation" && pgErr.Constraint == "users_uc_email" {
            	return ErrDuplicateEmail
        	}
    	}
    	return err
	}

	return nil
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

func (m *UserModel) Exists(id int) (bool, error) {
	return false, nil
}