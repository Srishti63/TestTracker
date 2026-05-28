package domain

import "context"

type Test struct {
	ID          string `json:"id" db:"_id"`
	TestGroupID string `json:"test_group_id" db:"test_group_id"`
	CreatedAt   string `json:"created_at" db:"created_at"` 
}


type Entry struct {
	ID              string  `json:"id" db:"_id"`
	TestID          string  `json:"test_id" db:"test_id"`
	SubjectPresetID string  `json:"subject_preset_id" db:"subject_preset_id"`
	MarksObtained   float64 `json:"marks_obtained" db:"marks_obtained"`
}


type TestRepository interface {
	CreateWithEntries(ctx context.Context, test *Test, entries []Entry) error
}

type TestUsecase interface {
	LogPerformance(ctx context.Context, test *Test, entries []Entry) error
}