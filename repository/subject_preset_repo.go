package repository

import (
	"context"
	"database/sql"
	"test_tracker_backend/domain"
	"github.com/google/uuid"
)

type subjectPresetRepository struct {
	db *sql.DB
}

func NewSubjectPresetRepository(db *sql.DB) domain.SubjectPresetRepository {
	return &subjectPresetRepository{db: db}
}

func (r *subjectPresetRepository) CreateMany(ctx context.Context, presets []domain.SubjectPreset) error {
	query := `INSERT INTO subject_presets (_id, test_group_id, subject_name, total_marks) VALUES ($1, $2, $3, $4);`
	
	for _, p := range presets {
		id := p.ID
		if id == "" {
			id = uuid.New().String()
		}
		_, err := r.db.ExecContext(ctx, query, id, p.TestGroupID, p.SubjectName, p.TotalMarks)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *subjectPresetRepository) GetByGroupID(ctx context.Context, groupID string) ([]domain.SubjectPreset, error) {
	query := `SELECT _id, test_group_id, subject_name, total_marks FROM subject_presets WHERE test_group_id = $1;`
	rows, err := r.db.QueryContext(ctx, query, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var presets []domain.SubjectPreset
	for rows.Next() {
		var p domain.SubjectPreset
		if err := rows.Scan(&p.ID, &p.TestGroupID, &p.SubjectName, &p.TotalMarks); err != nil {
			return nil, err
		}
		presets = append(presets, p)
	}
	return presets, nil
}

func (r *subjectPresetRepository) DeleteByGroupID(ctx context.Context, groupID string) error {
	query := `DELETE FROM subject_presets WHERE test_group_id = $1;`
	_, err := r.db.ExecContext(ctx, query, groupID)
	return err
}