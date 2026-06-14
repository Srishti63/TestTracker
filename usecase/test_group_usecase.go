package usecase

import (
	"context"
	"errors"
	"time"
	"test_tracker_backend/domain"
	"github.com/google/uuid"
)

type testGroupUsecase struct {
	repo           domain.TestGroupRepository
	contextTimeout time.Duration
}

func NewTestGroupUsecase(r domain.TestGroupRepository, timeout time.Duration) domain.TestGroupUsecase {
	return &testGroupUsecase{repo: r, contextTimeout: timeout}
}

func (u *testGroupUsecase) CreateGroup(ctx context.Context, req *domain.CreateTestGroupParam) (*domain.TestGroup,error) {
	c, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	if req.TestGroupName == "" || req.UserID == "" {
		return nil,errors.New("group name and credentials cannot be blank")
	}

	group := &domain.TestGroup{
		TestGroupID:   uuid.New().String(),
		UserID:        req.UserID,
		TestGroupName: req.TestGroupName,
	}

	err := u.repo.Create(c, group)
	if err != nil{
		return nil,err
	}

	return  group,nil
}

func (u *testGroupUsecase) FetchUserGroups(ctx context.Context, userID string) ([]domain.TestGroup, error) {
	c, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	if userID == "" {
		return nil, errors.New("missing user signature scope")
	}

	return u.repo.GetByByUserId(c, userID)
}