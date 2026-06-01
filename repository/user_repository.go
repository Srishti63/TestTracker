package repository

import (
	"context"
	"database/sql"
	"test_tracker_backend/domain"
	"time"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *domain.User) error {
	query := `
		INSERT INTO users (_id, name, email, passwordhash, created_at)
		VALUES ($1, $2, $3, $4, $5);`

	_, err := r.db.ExecContext(ctx, query, user.ID, user.Name, user.Email, user.PasswordHash, user.CreatedAt)
	return err
}

func (r *userRepository) GetByName(ctx context.Context, name string) (*domain.User, error) {
	query := `
		SELECT _id, name, passwordhash, created_at, email FROM users WHERE name = $1;`

	row := r.db.QueryRowContext(ctx, query, name)

	var user domain.User
	err := row.Scan(&user.ID, &user.Name, &user.PasswordHash, &user.CreatedAt, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByID(ctx context.Context, ID string) (*domain.User, error) {
	query := `SELECT _id, name, passwordhash, created_at, email FROM users WHERE _id = $1;`

	row := r.db.QueryRowContext(ctx, query, ID)

	var user domain.User
	err := row.Scan(&user.ID, &user.Name, &user.PasswordHash, &user.CreatedAt, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `SELECT _id, name, passwordhash, created_at, email FROM users WHERE email = $1;`

	row := r.db.QueryRowContext(ctx, query, email)

	var user domain.User
	err := row.Scan(&user.ID, &user.Name, &user.PasswordHash, &user.CreatedAt, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil 
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdatePassword(ctx context.Context, userID string, hashedNewPassword string) error {
	query := `
		UPDATE users 
		SET passwordhash = $1
		WHERE _id = $2;`

	// Passed BOTH variables in the exact placeholder order: $1, $2
	_, err := r.db.ExecContext(ctx, query, hashedNewPassword, userID)
	return err
}

func (r *userRepository) UpdateResetToken(ctx context.Context, userID string, token string, expiry time.Time) error {
	query := `
		UPDATE users 
		SET reset_token = $1, reset_token_expiry = $2 
		WHERE _id = $3;`

	// Pass variables in exact matching placeholder sequence: $1, $2, $3
	_, err := r.db.ExecContext(ctx, query, token, expiry, userID)
	return err
}

func (r *userRepository) ClearResetToken(ctx context.Context, userID string) error {
	query := `
		UPDATE users 
		SET reset_token = NULL, reset_token_expiry = NULL 
		WHERE _id = $1;`

	_, err := r.db.ExecContext(ctx, query, userID)
	return err
}