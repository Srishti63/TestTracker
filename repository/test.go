package repository

import (
	"context"
	"database/sql"
	"test_tracker_backend/domain"
	"github.com/google/uuid"
)

type testRepository struct {
	db *sql.DB
}

func NewTestRepository(db *sql.DB) domain.TestRepository {
	return &testRepository{db: db}
}

func (r *testRepository) CreateWithEntries(ctx context.Context, test *domain.Test, entries []domain.Entry) error {
	// 1. Begin a secure database transaction
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	// Defer an automatic rollback. If the transaction commits successfully, this becomes a no-op.
	defer tx.Rollback()

	// 2. Drop the parent row into the tests table
	testQuery := `INSERT INTO tests (_id, test_group_id, created_at) VALUES ($1, $2, $3);`
	_, err = tx.ExecContext(ctx, testQuery, test.ID, test.TestGroupID, test.CreatedAt)
	if err != nil {
		return err
	}

	// 3. Loop over all subject score entries and insert them using the transaction pointer
	entryQuery := `INSERT INTO entries (_id, test_id, subject_preset_id, marks_obtained) VALUES ($1, $2, $3, $4);`
	for _, entry := range entries {
		entryID := uuid.New().String()
		_, err = tx.ExecContext(ctx, entryQuery, entryID, test.ID, entry.SubjectPresetID, entry.MarksObtained)
		if err != nil {
			return err // Triggers automatic rollback of everything, including the test row!
		}
	}

	// 4. Explicitly commit the changes if everything succeeded
	return tx.Commit()
}