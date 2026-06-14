package dto

import "time"

type ScoreEntryItem struct {
	SubjectID     string  `json:"subject_id" binding:"required,uuid4"`
	MarksObtained float64 `json:"marks_obtained" binding:"required,gte=0"`
}

type CreateTestRequest struct {
	TestGroupID string           `json:"test_group_id" binding:"required,uuid4"`
	Title       string           `json:"title" binding:"required,min=2,max=100"`
	Scores      []ScoreEntryItem `json:"scores" binding:"required,dive"`
}

type TestEntryResponse struct {
	SubjectID     string  `json:"subject_id"`
	SubjectName   string  `json:"subject_name"`
	MarksObtained float64 `json:"marks_obtained"`
}

type TestResponse struct {
	ID          string              `json:"id"`
	TestGroupID string              `json:"test_group_id"`
	Title       string              `json:"title"`
	CreatedAt   time.Time           `json:"created_at"`
	Scores      []TestEntryResponse `json:"scores"`
}