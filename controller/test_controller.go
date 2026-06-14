package controller

import (
	"net/http"
	"test_tracker_backend/domain"
	"test_tracker_backend/dto"
	"github.com/gin-gonic/gin"
)

type testController struct {
	usecase domain.TestUsecase
}

func NewTestController(u domain.TestUsecase) *testController {
	return &testController{usecase: u}
}

func (t *testController) Create(c *gin.Context) {
	ctx := c.Request.Context()

	var req dto.CreateTestRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid test session layout"})
		return
	}

	var entryParams []domain.CreateEntryParam
	for _, s := range req.Scores {
		entryParams = append(entryParams, domain.CreateEntryParam{
			SubjectID:     s.SubjectID,
			MarksObtained: s.MarksObtained,
		})
	}

	param := &domain.CreateTestParam{
		TestGroupID: req.TestGroupID,
		Title:       req.Title,
		Entries:     entryParams,
	}

	if err := t.usecase.CreateTest(ctx, param); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Test score record logged successfully"})
}

func (t *testController) FetchByGroup(c *gin.Context) {
	ctx := c.Request.Context()
	groupID := c.Param("id")

	tests, err := t.usecase.FetchGroupTests(ctx, groupID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var responseData []dto.TestResponse
	for _, test := range tests {
		var scoreEntries []dto.TestEntryResponse
		for _, e := range test.Entries {
			scoreEntries = append(scoreEntries, dto.TestEntryResponse{
				SubjectID:     e.SubjectID,
				SubjectName:   e.SubjectName,
				MarksObtained: e.MarksObtained,
			})
		}

		responseData = append(responseData, dto.TestResponse{
			ID:          test.ID,
			TestGroupID: test.TestGroupID,
			Title:       test.Title,
			CreatedAt:   test.CreatedAt,
			Scores:      scoreEntries,
		})
	}

	c.JSON(http.StatusOK, responseData)
}