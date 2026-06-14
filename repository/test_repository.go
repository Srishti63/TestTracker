package repository

import (
	"context"
	"database/sql"
	"test_tracker_backend/domain"
)

type testRepository struct {
	db *sql.DB
}

func NewTestRepository(db *sql.DB) domain.TestRepository {
	return &testRepository{db: db}
}

func (r *testRepository) Create(ctx context.Context, test *domain.Test) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// 1. Insert into tests
	testQuery := `INSERT INTO tests (id, test_group_id, title, created_at) VALUES ($1, $2, $3, $4)`
	_, err = tx.ExecContext(ctx, testQuery, test.ID, test.TestGroupID, test.Title, test.CreatedAt)
	if err != nil {
		return err
	}

	// 2. Insert into test_entries
	entryQuery := `INSERT INTO test_entries (id, test_id, subject_id, marks_obtained) VALUES ($1, $2, $3, $4)`
	for _, e := range test.Entries {
		_, err = tx.ExecContext(ctx, entryQuery, e.ID, e.TestID, e.SubjectID, e.MarksObtained)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *testRepository) GetByGroupID(ctx context.Context, groupID string) ([]domain.Test, error) {
	// First, get all tests in the group
	testQuery := `SELECT id, test_group_id, title, created_at FROM tests WHERE test_group_id = $1 ORDER BY created_at ASC`
	testRows, err := r.db.QueryContext(ctx, testQuery, groupID)
	if err != nil {
		return nil, err
	}
	defer testRows.Close()

	var tests []domain.Test
	for testRows.Next() {
		var t domain.Test
		if err := testRows.Scan(&t.ID, &t.TestGroupID, &t.Title, &t.CreatedAt); err != nil {
			return nil, err
		}
		tests = append(tests)
	}

	// For each test, fetch its inner subject scores using a SQL JOIN
	for i := range tests {
		entryQuery := `
			SELECT te.id, te.test_id, te.subject_id, s.name as subject_name, te.marks_obtained 
			FROM test_entries te
			JOIN subjects s ON te.subject_id = s.id
			WHERE te.test_id = $1`
		
		entryRows, err := r.db.QueryContext(ctx, entryQuery, tests[i].ID)
		if err != nil {
			return nil, err
		}
		
		var entries []domain.TestEntry
		for entryRows.Next() {
			var e domain.TestEntry
			if err := entryRows.Scan(&e.ID, &e.TestID, &e.SubjectID, &e.SubjectName, &e.MarksObtained); err != nil {
				entryRows.Close()
				return nil, err
			}
			entries = append(entries, e)
		}
		entryRows.Close()
		tests[i].Entries = entries
	}

	return tests, nil
}