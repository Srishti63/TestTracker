package route

import (
	"test_tracker_backend/controller"
	"test_tracker_backend/domain"
	"test_tracker_backend/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func Setup (r *gin.Engine , userRepo domain.UserRepository, timeOut time.Duration , jwtSecret string, jwtExpiryhrs int){
	
	userUsecase := usecase.NewUserUsecase(userRepo,timeOut,jwtSecret,jwtExpiryhrs)
	userController := controller.NewUserController(userUsecase)

	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/register", userController.Register)
			auth.POST("/login", userController.Login)
			auth.POST("/forgot-password", userController.ForgotPassword)
			auth.POST("/confirm-password", userController.ConfirmPasswordReset)
		}

		tests := v1.Group("/tests")
		{
			// Fetch presets for a specific test group to build the frontend form
			tests.GET("/presets/:group_id", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Placeholder: Fetching presets blueprint"})
			})
			
			// Submit/Log the students performance marks payload
			tests.POST("/log", func(c *gin.Context) {
				c.JSON(201, gin.H{"message": "Placeholder: Logging performance blueprint"})
			})
		}
	}
}