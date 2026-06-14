package controller

import (
	"net/http"
	"test_tracker_backend/domain"
	"test_tracker_backend/dto"
	"github.com/gin-gonic/gin"
)

type subjectPresetController struct {
	Usecase domain.SubjectPresetUsecase
}

func NewSubjectPresetController(u domain.SubjectPresetUsecase) *subjectPresetController {
	return &subjectPresetController{Usecase: u}
}

func (p *subjectPresetController) BatchConfigure(c *gin.Context) {
	ctx := c.Request.Context()

	var req dto.BatchPresetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid dynamic schema batch body layout"})
		return
	}

	var corePresets []domain.SubjectPreset
	for _, item := range req.Presets {
		corePresets = append(corePresets, domain.SubjectPreset{
			TestGroupID: req.TestGroupID,
			SubjectID:   item.SubjectID,
			TotalMarks:  item.TotalMarks,
		})
	}

	if err := p.Usecase.ConfigurePresets(ctx, corePresets); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Presets configuration mapped safely"})
}

func (p *subjectPresetController) FetchByGroup(c *gin.Context) {
	ctx := c.Request.Context()
	testGroupID := c.Param("id")

	presets, err := p.Usecase.FetchGroupPresets(ctx, testGroupID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var responseData []dto.PresetConfigResponse
	for _, pr := range presets {
		responseData = append(responseData, dto.PresetConfigResponse{
			ID:          pr.ID,
			TestGroupID: pr.TestGroupID,
			SubjectID:   pr.SubjectID,
			SubjectName: pr.SubjectName,
			TotalMarks:  pr.TotalMarks,
		})
	}

	c.JSON(http.StatusOK, responseData)
}