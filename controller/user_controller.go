package controller

import (
	"net/http"
	"test_tracker_backend/domain"

	"github.com/gin-gonic/gin"
)

type userController struct {
	UserUsecase domain.UserUsecase
}

func NewUserController(usecase domain.UserUsecase) domain.UserController {
	return &userController{
		UserUsecase: usecase,
	}
}


func (u *userController) Register(c *gin.Context) {
	ctx := c.Request.Context()
	var dto RegisterDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	domainReq := dto.ToDomain()
	err := u.UserUsecase.Register(ctx, domainReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully!"})
}

func (u *userController) Login(c *gin.Context) {
	ctx := c.Request.Context()
	var dto LoginDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	domainReq := dto.ToDomain()
	token, err := u.UserUsecase.Login(ctx, domainReq)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Login successful!",
		"access_token": token,
	})
}


func (u *userController) ForgotPassword(c *gin.Context) {
	ctx := c.Request.Context()
	var dto ForgotPasswordDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	domainReq := dto.ToDomain()
	err := u.UserUsecase.ForgotPassword(ctx, domainReq.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset token sent to your email!"})
}


func (u *userController) ConfirmPasswordReset(c *gin.Context) {
	ctx := c.Request.Context()
	var dto ConfirmPasswordResetDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	domainReq := dto.ToDomain()
	err := u.UserUsecase.ConfirmPasswordReset(ctx, domainReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully!"})
}