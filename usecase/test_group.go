package usecase

import (
	"context"
	"time"
	"errors"
	"test_tracker_backend/domain"
)

type testGroupUsecase struct {
	testGroupRepo  domain.TestGroupRepository
	contextTimeout time.Duration
}

// NewTestGroupUsecase injects the matching repository interface contract dependency
func NewTestGroupUsecase(repo domain.TestGroupRepository, timeout time.Duration) domain.TestGroupUsecase {
	return &testGroupUsecase{
		testGroupRepo:  repo,
		contextTimeout: timeout,
	}
}

// CreateGroup manages constraints or transformations before execution injection
func (u *testGroupUsecase) CreateGroup(ctx context.Context, group *domain.TestGroup) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	// Business rule validation constraint check
	if group.TestGroupName == "" {
		return errors.New("can't fetch the testGroup name")
	}

	return u.testGroupRepo.Create(ctx, group)
}

// FetchUserGroups acts as the clean pipe coordinator to read historical models
func (u *testGroupUsecase) FetchUserGroups(ctx context.Context, userID string) ([]domain.TestGroup, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.testGroupRepo.GetByUserId(ctx, userID)
}