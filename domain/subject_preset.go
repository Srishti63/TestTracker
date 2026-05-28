package domain

import "context"

type SubjectPreset struct {
	ID          string  `json:"id" db:"_id"`
	TestGroupID string  `json:"test_group_id" db:"test_group_id"`
	SubjectName string  `json:"subject_name" db:"subject_name"`
	TotalMarks  float64 `json:"total_marks" db:"total_marks"`
}

type SubjectPresetRepository interface {
	CreateMany(ctx context.Context, presets []SubjectPreset) error
	GetByGroupID(ctx context.Context, groupID string) ([]SubjectPreset, error)
}

type SubjectPresetUsecase interface {
	ConfigurePresets(ctx context.Context, presets []SubjectPreset) error
	FetchGroupPresets(ctx context.Context, groupID string) ([]SubjectPreset, error)
}