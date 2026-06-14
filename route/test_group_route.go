package route

import (
	"test_tracker_backend/controller"
	"test_tracker_backend/domain"
	"github.com/gin-gonic/gin"
)

func SetupTestGroupRoutes(r *gin.Engine, testgroupUsecase domain.TestGroupUsecase) {
	ctrl := controller.NewTestGroupController(testgroupUsecase)

	// In Hoppscotch, make sure your JWT Auth Middleware intercepts this group!
	api := r.Group("/api/v1/test-groups")
	api.Use(func(c *gin.Context) {
        c.Set("x-user-id", "f90597ec-aafb-4e88-b481-e2e1fef74668") // Same DB user UUID
        c.Next()
    })
	{
		api.POST("", ctrl.Create)
		api.GET("", ctrl.FetchAll)
	}
}