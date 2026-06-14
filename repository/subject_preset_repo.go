package repository

import (
	"context"
	"database/sql"
	"test_tracker_backend/domain"
)

type subjectPresetRepository struct {
	db *sql.DB
}

func NewSubjectPresetRepository(db *sql.DB) domain.SubjectPresetRepository {
	return &subjectPresetRepository{db: db}
}

func (r *subjectPresetRepository) CreateMany(ctx context.Context, presets []domain.SubjectPreset) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `INSERT INTO subject_presets (id, test_group_id, subject_id, total_marks) VALUES ($1, $2, $3, $4)`
	for _, p := range presets {
		if _, err := tx.ExecContext(ctx, query, p.ID, p.TestGroupID, p.SubjectID, p.TotalMarks); err != nil {
			return err
		}
	}
	return tx.Commit()
}

func (r *subjectPresetRepository) GetByGroupID(ctx context.Context, groupID string) ([]domain.SubjectPreset, error) {
	query := `
		SELECT sp.id, sp.test_group_id, sp.subject_id, s.name as subject_name, sp.total_marks 
		FROM subject_presets sp
		JOIN subjects s ON sp.subject_id = s.id
		WHERE sp.test_group_id = $1`

	rows, err := r.db.QueryContext(ctx, query, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var presets []domain.SubjectPreset
	for rows.Next() {
		var p domain.SubjectPreset
		if err := rows.Scan(&p.ID, &p.TestGroupID, &p.SubjectID, &p.SubjectName, &p.TotalMarks); err != nil {
			return nil, err
		}
		presets = append(presets, p)
	}
	return presets, nil
}