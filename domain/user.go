package domain

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

// User represents the core database entity schema mapped to PostgreSQL.
type User struct {
	ID                 string     `json:"id" db:"_id"`
	Email              string     `json:"email" db:"email"`
	PasswordHash       string     `json:"-" db:"passwordhash"` // "-" prevents hash from leaking in JSON marshaling
	Name               string     `json:"name" db:"name"`
	CreatedAt          time.Time  `json:"created_at" db:"created_at"`
	ResetToken         *string    `json:"-" db:"reset_token"`        // Pointer type allows NULL in DB
	ResetTokenExpiry   *time.Time `json:"-" db:"reset_token_expiry"` // Pointer type allows NULL in DB
}

// --- Request DTOs (Data Transfer Objects) for Controllers ---

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResetPasswordRequest struct {
	Email       string `json:"email"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type ConfirmPasswordResetRequest struct {
	Email       string `json:"email" binding:"required,email"`
	RandomToken string `json:"random_token" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}

type UpdatePasswordRequest struct {
	UserID      string `json:"user_id" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}

type ForgotPasswordRequest struct{
	Email string `json:"email" binding:"required,email"`

}
type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id string) (*User, error)
	GetByName(ctx context.Context, name string) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	UpdatePassword(ctx context.Context, userID string, hashedNewPassword string) error
	UpdateResetToken(ctx context.Context, userID string, token string, expiry time.Time) error
	ClearResetToken(ctx context.Context, userID string) error
}

type UserUsecase interface {
	Register(ctx context.Context, req *RegisterRequest) error
	Login(ctx context.Context, req *LoginRequest) (string, error) // Returns signed JWT token string
	ResetPassword(ctx context.Context, req *ResetPasswordRequest) error
	UpdatePassword(ctx context.Context, userID string, newPassword string) error
	ForgotPassword(ctx context.Context, email string) error
	ConfirmPasswordReset(ctx context.Context, req *ConfirmPasswordResetRequest) error
}

type UserController interface {
    Register(c *gin.Context)
    Login(c *gin.Context)
    ForgotPassword(c *gin.Context)
    ConfirmPasswordReset(c *gin.Context)
}