package controller

import (
	"net/http"
	"test_tracker_backend/domain"
	"github.com/gin-gonic/gin"
	"test_tracker_backend/dto"
)

type subjectController struct {
	SubjectUsecase domain.SubjectUsecase
}

func NewSubjectController(su domain.SubjectUsecase) *subjectController {
	return &subjectController{
		SubjectUsecase: su,
	}
}

func (s *subjectController) Create(c *gin.Context) {
	ctx := c.Request.Context()
	userID, _ := c.Get("x-user-id")

	var req dto.CreateSubjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Validation constraints failed"})
		return
	}

	param := &domain.CreateSubjectParam{
		Name:   req.Name,
		UserID: userID.(string),
	}

	createdSubject,err := s.SubjectUsecase.CreateSubject(ctx, param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := dto.SubjectResponse{
		ID: createdSubject.ID,
		Name: createdSubject.Name,
	}

	c.JSON(http.StatusCreated,res)
}


func (s *subjectController) FetchAll(c *gin.Context) {
	ctx := c.Request.Context()

	
	userID, exists := c.Get("x-user-id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication context missing or expired"})
		return
	}

	domainSubjects, err := s.SubjectUsecase.FetchUserSubjects(ctx, userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sync catalog items"})
		return
	}

	var responseData []dto.SubjectResponse
	for _, sub := range domainSubjects {
		responseData = append(responseData, dto.SubjectResponse{
			ID:   sub.ID,
			Name: sub.Name,
		})
	}

	c.JSON(http.StatusOK, responseData)
}