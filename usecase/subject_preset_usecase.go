package usecase

import (
	"context"
	"errors"
	"time"
	"test_tracker_backend/domain"
	"github.com/google/uuid"
)

type subjectPresetUsecase struct {
	repo           domain.SubjectPresetRepository
	contextTimeout time.Duration
}

func NewSubjectPresetUsecase(r domain.SubjectPresetRepository, timeout time.Duration) domain.SubjectPresetUsecase {
	return &subjectPresetUsecase{repo: r, contextTimeout: timeout}
}

func (u *subjectPresetUsecase) ConfigurePresets(ctx context.Context, presets []domain.SubjectPreset) error {
	c, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	if len(presets) == 0 {
		return errors.New("presets matrix configuration cannot be empty")
	}

	for i := range presets {
		presets[i].ID = uuid.New().String()
	}

	return u.repo.CreateMany(c, presets)
}

func (u *subjectPresetUsecase) FetchGroupPresets(ctx context.Context, groupID string) ([]domain.SubjectPreset, error) {
	c, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	if groupID == "" {
		return nil, errors.New("empty test group tracker scope mapping pointer")
	}

	return u.repo.GetByGroupID(c, groupID)
}