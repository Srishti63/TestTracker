package controller

import (
	"net/http"
	"test_tracker_backend/domain"
	"test_tracker_backend/usecase"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/message"
)
 
type userController struct {
	UserUsecase domain.UserUsecase
}

func NewUserController (usecase domain.UserUsecase) *userController{
	return &userController{
		UserUsecase: usecase,
	}
}

func (u *userController) Login ( c *gin.Context){
	ctx := c.Request.Context()

	var req domain.LoginRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":  "Invalid request body or missing credentials"})
		return
	}

	token, err := u.UserUsecase.Login(ctx, &req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error" : err.Error()})
		return 
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "Login successful",
		"access-token" : token,
	})
}

func (u *userController) Register(c *gin.Context){
	ctx := c. Request.Context()

	var req domain.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	err := u.UserUsecase.Register(ctx, &req)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return 
	}

	c.JSON(http.StatusCreated, gin.H{
		"message" : "Registered successfully",
	})
}

func (u *userController) ForgotPassword (c *gin.Context){
	ctx := c.Request.Context()

	var req domain.ForgotPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"Error": "A valid email address is required"})
		return 
	}

	err := u.UserUsecase.ForgotPassword(ctx,req.Email)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal communication channel breakdown"})
	return
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "If the account exists , a secure verification has been dispatched",
	})
}

func (u *userController) ConfirmPasswordReset(c *gin.Context){
	ctx := c.Request.Context()

	var req domain.ConfirmPasswordResetRequest
	if err := c.ShouldBindJSON(&req); err != nil{
		
	}

}