package repositories

import (
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func (r *UserRepository) FindByID(id string) (*User, error)