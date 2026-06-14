package usecase

import (
	"context"
	"errors"
	"time"
	"test_tracker_backend/domain"
	"github.com/google/uuid"
)

type testUsecase struct {
	repo           domain.TestRepository
	contextTimeout time.Duration
}

func NewTestUsecase(r domain.TestRepository, timeout time.Duration) domain.TestUsecase {
	return &testUsecase{repo: r, contextTimeout: timeout}
}

func (u *testUsecase) CreateTest(ctx context.Context, param *domain.CreateTestParam) error {
	c, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	if param.Title == "" || param.TestGroupID == "" || len(param.Entries) == 0 {
		return errors.New("incomplete test parameters; scores cannot be empty")
	}

	testID := uuid.New().String()
	var coreEntries []domain.TestEntry

	for _, e := range param.Entries {
		coreEntries = append(coreEntries, domain.TestEntry{
			ID:            uuid.New().String(),
			TestID:        testID,
			SubjectID:     e.SubjectID,
			MarksObtained: e.MarksObtained,
		})
	}

	test := &domain.Test{
		ID:          testID,
		TestGroupID: param.TestGroupID,
		Title:       param.Title,
		CreatedAt:   time.Now(),
		Entries:     coreEntries,
	}

	return u.repo.Create(c, test)
}

func (u *testUsecase) FetchGroupTests(ctx context.Context, groupID string) ([]domain.Test, error) {
	c, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	if groupID == "" {
		return nil, errors.New("test group tracking container reference missing")
	}

	return u.repo.GetByGroupID(c, groupID)
}