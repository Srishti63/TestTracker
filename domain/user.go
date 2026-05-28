package domain

import (
	"context"
	"time"
)

type User struct {
	ID           string    `json:"id" db:"_id"`
	Name         string    `json:"name" db:"name"`
	PasswordHash string    `json:"-" db:"password_hash"` 
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByName(ctx context.Context, name string) (*User, error)
	GetByID(ctx context.Context, id string) (*User, error)
}

type SignupUsecase interface {
	Register(ctx context.Context, user *User, plainPassword string) error
}