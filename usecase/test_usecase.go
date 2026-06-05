package usecase

import(
	"time"
	"context"
	"errors"
	"test_tracker_backend/domain"
	"github.com/google/uuid"
)

type testUsecase struct{
	testRepo domain.TestRepository
	contextTimeout time.Duration
}

func NewTestUsecase(repo domain.TestRepository , timeout time.Duration) domain.TestUsecase{
	return &testUsecase{
		testRepo : repo,
		contextTimeout: timeout,
	}
}

func (tu* testUsecase) LogPerformance(ctx context.Context, test *domain.Test , entries []domain.Entry) error {
	ctx , cancel := context.WithTimeout(ctx , tu.contextTimeout)
	defer cancel()

	if test.TestGroupID == "" {
		return errors.New("cannot log a test without a valid test group identifier")
	}
	if len(entries) == 0 {
		return errors.New("cannot log a test event with zero subject scores")
	}

	test.ID = uuid.New().String()

	test.CreatedAt = time.Now().Format("2006-01-02")

	// 3. Link the child score slices back to this parent Test record
	for i := range entries {
		if entries[i].MarksObtained < 0 {
			return errors.New("marks obtained cannot be negative values")
		}
		// Essential step: assign the parent foreign key link
		entries[i].TestID = test.ID
	}

	// 4. Pass the fully assembled objects down to the repository database layer
	// This invokes your database transaction block to write rows safely!
	return tu.testRepo.CreateWithEntries(ctx, test, entries)

}