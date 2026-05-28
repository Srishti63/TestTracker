package usecase

import (
	"context"
	"errors"
	"time"
	"test_tracker_backend/domain"
)

type subjectPresetUsecase struct {
	presetRepo     domain.SubjectPresetRepository
	contextTimeout time.Duration
}

func NewSubjectPresetUsecase(repo domain.SubjectPresetRepository, timeout time.Duration) domain.SubjectPresetUsecase {
	return &subjectPresetUsecase{
		presetRepo:     repo,
		contextTimeout: timeout,
	}
}

func (u *subjectPresetUsecase) ConfigurePresets(ctx context.Context, presets []domain.SubjectPreset) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	if len(presets) == 0 {
		return errors.New("cannot configure an empty set of subjects")
	}

	for _, p := range presets {
		if p.SubjectName == "" || p.TotalMarks <= 0 {
			return errors.New("invalid subject criteria: name required and marks must be greater than 0")
		}
	}

	return u.presetRepo.CreateMany(ctx, presets)
}

func (u *subjectPresetUsecase) FetchGroupPresets(ctx context.Context, groupID string) ([]domain.SubjectPreset, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	if groupID == "" {
		return nil, errors.New("group id parameter is mandatory")
	}

	return u.presetRepo.GetByGroupID(ctx, groupID)
}