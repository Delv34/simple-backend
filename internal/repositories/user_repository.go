package repositories

import (
	"context"
	"database/sql"
	"go-simple-backend/internal/models"
	"time"
)

type UserRepository interface {
	Create(ctx context.Context, username, password string) (*models.User, error)
	FindByID(ctx context.Context, id uint32) (*models.User, error)
	FindAll(ctx context.Context) ([]*models.User, error)
	Update(ctx context.Context, id uint32, username, password *string) (*models.User, error)
	Delete(ctx context.Context, id uint32) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, username, password string) (*models.User, error) {
	createdAt := time.Now()
	result, err := r.db.ExecContext(ctx, `INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
	if err != nil {
		return nil, err
	}
	id, _ := result.LastInsertId()
	return &models.User{
		ID:        uint32(id),
		Username:  username,
		Password:  password,
		CreatedAt: createdAt,
	}, nil
}

func (r *userRepository) FindByID(ctx context.Context, id uint32) (*models.User, error) {
	var u models.User
	err := r.db.QueryRowContext(ctx, `SELECT id, username, created_at FROM users WHERE id = ?`, id).Scan(&u.ID, &u.Username, &u.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // можно написать user not found
		}
		return nil, err
	}
	return &u, nil
}

func (r *userRepository) FindAll(ctx context.Context) ([]*models.User, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, username, created_at FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Username, &u.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, &u)
	}
	return users, rows.Err()
}

func (r *userRepository) Update(ctx context.Context, id uint32, username, password *string) (*models.User, error) {
	query := `UPDATE users SET `
	args := []interface{}{}
	if username != nil {
		query += `username = ?, `
		args = append(args, *username)
	}
	if password != nil {
		query += `password = ?, `
		args = append(args, *password)
	}
	query = query[:len(query)-2] + ` WHERE id = ?`
	args = append(args, id)

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return r.FindByID(ctx, id)
}

func (r *userRepository) Delete(ctx context.Context, id uint32) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM users WHERE id = ?`, id)
	return err
}
