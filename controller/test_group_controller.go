package controller

import (
	"net/http"
	"test_tracker_backend/domain"
	"test_tracker_backend/dto"
	"github.com/gin-gonic/gin"
)

type testGroupController struct {
	Usecase domain.TestGroupUsecase
}

func NewTestGroupController(u domain.TestGroupUsecase) *testGroupController {
	return &testGroupController{Usecase: u}
}

func (t *testGroupController) Create(c *gin.Context) {
	ctx := c.Request.Context()
	userID, _ := c.Get("x-user-id")

	var req dto.CreateTestGroupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid test group title layout"})
		return
	}

	param := &domain.CreateTestGroupParam{
		TestGroupName: req.TestGroupName,
		UserID:        userID.(string),
	}

	// 1. Call the usecase and unpack the variables cleanly
    createdGroup, err := t.Usecase.CreateGroup(ctx, param)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) // No comma
        return // No comma
    }

    // 2. Map to your response DTO
    res := dto.TestGroupResponse{
        TestGroupID:            createdGroup.TestGroupID,
        TestGroupName: createdGroup.TestGroupName,
    }

    c.JSON(http.StatusCreated, res)
}

func (t *testGroupController) FetchAll(c *gin.Context) {
	ctx := c.Request.Context()
	userID, _ := c.Get("x-user-id")

	groups, err := t.Usecase.FetchUserGroups(ctx, userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load group metrics"})
		return
	}

	var responseData []dto.TestGroupResponse
	for _, g := range groups {
		responseData = append(responseData, dto.TestGroupResponse{
			TestGroupID:   g.TestGroupID,
			TestGroupName: g.TestGroupName,
		})
	}

	c.JSON(http.StatusOK, responseData)
}