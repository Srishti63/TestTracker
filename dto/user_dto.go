package dto

import "test_tracker_backend/domain"

// ==========================================
// REGISTER DTO
// ==========================================
type RegisterDTO struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (d *RegisterDTO) ToDomain() *domain.RegisterRequest {
	return &domain.RegisterRequest{
		Name:     d.Name,
		Email:    d.Email,
		Password: d.Password,
	}
}

// ==========================================
// LOGIN DTO
// ==========================================
type LoginDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (d *LoginDTO) ToDomain() *domain.LoginRequest {
	return &domain.LoginRequest{
		Email:    d.Email,
		Password: d.Password,
	}
}

// ==========================================
// FORGOT PASSWORD DTO
// ==========================================
type ForgotPasswordDTO struct {
	Email string `json:"email" binding:"required,email"`
}

func (d *ForgotPasswordDTO) ToDomain() *domain.ForgotPasswordRequest {
	return &domain.ForgotPasswordRequest{
		Email: d.Email,
	}
}

// ==========================================
// CONFIRM PASSWORD RESET DTO
// ==========================================
type ConfirmPasswordResetDTO struct {
	Email       string `json:"email" binding:"required,email"`
	RandomToken string `json:"token" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}

func (d *ConfirmPasswordResetDTO) ToDomain() *domain.ConfirmPasswordResetRequest {
	return &domain.ConfirmPasswordResetRequest{
		Email:       d.Email,
		RandomToken: d.RandomToken,
		NewPassword: d.NewPassword,
	}
}