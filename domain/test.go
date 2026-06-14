package domain

import (
	"context"
	"time"
)

type Test struct {
	ID          string    `json:"id" db:"id"`
	TestGroupID string    `json:"test_group_id" db:"test_group_id"`
	Title       string    `json:"title" db:"title"` // e.g., "Mock Test 1"
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	Entries     []TestEntry `json:"entries,omitempty"` // Nested list of scores
}

type TestEntry struct {
	ID        string  `json:"id" db:"id"`
	TestID    string  `json:"test_id" db:"test_id"`
	SubjectID string  `json:"subject_id" db:"subject_id"`
	SubjectName string `json:"subject_name,omitempty" db:"subject_name"` // From SQL JOIN
	MarksObtained float64 `json:"marks_obtained" db:"marks_obtained"`
}

type CreateTestParam struct {
	TestGroupID string
	Title       string
	Entries     []CreateEntryParam
}

type CreateEntryParam struct {
	SubjectID     string
	MarksObtained float64
}

type TestRepository interface {
	Create(ctx context.Context, test *Test) error
	GetByGroupID(ctx context.Context, groupID string) ([]Test, error)
}

type TestUsecase interface {
	CreateTest(ctx context.Context, param *CreateTestParam) error
	FetchGroupTests(ctx context.Context, groupID string) ([]Test, error)
}