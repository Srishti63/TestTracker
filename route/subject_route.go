package route

import (
	"test_tracker_backend/controller"
	"test_tracker_backend/domain"

	"github.com/gin-gonic/gin"
)

func SetupSubjectRoutes(r *gin.Engine, subjetUsecase domain.SubjectUsecase) {
	ctrl := controller.NewSubjectController(subjetUsecase)


	api := r.Group("/api/v1/subjects")
	
	// temporarily skipping the need to write massive middleware code , will add it up at night 
	api.Use(func(c *gin.Context) {
        c.Set("x-user-id", "f90597ec-aafb-4e88-b481-e2e1fef74668") // Use your real DB user UUID here!
        c.Next()
    })

	{
		api.POST("", ctrl.Create)
		api.GET("", ctrl.FetchAll)
	}
}