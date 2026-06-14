package route

import (
	"test_tracker_backend/controller"
	"test_tracker_backend/domain"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine, uc domain.UserUsecase) {
	ctrl := controller.NewUserController(uc)

	api := r.Group("/api/v1/auth")
	{
		api.POST("/register", ctrl.Register)
		api.POST("/login", ctrl.Login)
		api.POST("/forgot-password", ctrl.ForgotPassword)
		api.POST("/reset-password", ctrl.ConfirmPasswordReset)
	}
}