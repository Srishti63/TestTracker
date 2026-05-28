package repository

import (
	"context"
	"database/sql"
	"test_tracker_backend/domain"
)

type testGroupRepository struct {
	db *sql.DB
}

// NewTestGroupRepository hooks up the database engine
func NewTestGroupRepository(db *sql.DB) domain.TestGroupRepository {
	return &testGroupRepository{db: db}
}

// Create runs the SQL INSERT command to save the new tracking layout metadata
func (r *testGroupRepository) Create(ctx context.Context, group *domain.TestGroup) error {
	query := `
		INSERT INTO test_groups (test_group_id, user_id, test_group_name) 
		VALUES ($1, $2, $3);`

	_, err := r.db.ExecContext(ctx, query, group.TestGroupId, group.UserID, group.TestGroupName)
	return err
}

// GetByUserID pulls down every configured template class belonging to the user
func (r *testGroupRepository) GetByUserId(ctx context.Context, userID string) ([]domain.TestGroup, error) {
	query := `
		SELECT test_group_id, user_id, test_group_name 
		FROM test_groups 
		WHERE user_id = $1;`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []domain.TestGroup
	for rows.Next() {
		var g domain.TestGroup
		if err := rows.Scan(&g.TestGroupId, &g.UserID, &g.TestGroupName); err != nil {
			return nil, err
		}
		groups = append(groups, g)
	}
	
	return groups, nil
}