package route

import (
	"test_tracker_backend/controller"
	"test_tracker_backend/domain"
	"github.com/gin-gonic/gin"
)

func SetupSubjectPresetRoutes(r *gin.Engine, subjectPreset domain.SubjectPresetUsecase) {
	ctrl := controller.NewSubjectPresetController(subjectPreset)

	api := r.Group("/api/v1/presets")
	{
		api.POST("/batch", ctrl.BatchConfigure)
		api.GET("/group/:id", ctrl.FetchByGroup)
	}
}