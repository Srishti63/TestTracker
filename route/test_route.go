package route

import (
	"test_tracker_backend/controller"
	"test_tracker_backend/domain"
	"github.com/gin-gonic/gin"
)

func SetupTestRoutes(r *gin.Engine, testUsecase domain.TestUsecase) {
	ctrl := controller.NewTestController(testUsecase)

	api := r.Group("/api/v1/tests")
	{
		api.POST("", ctrl.Create)
		api.GET("/group/:id", ctrl.FetchByGroup) // 📈 This endpoint feeds your React Performance Chart!
	}
}