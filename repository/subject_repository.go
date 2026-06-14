package repository

import (
	"context"
	"database/sql"
	"test_tracker_backend/domain"
)

type subjectRepository struct {
	db *sql.DB
}

func NewSubjectRepository(db *sql.DB) domain.SubjectRepository {
	return &subjectRepository{db: db}
}

func (r *subjectRepository) Create(ctx context.Context, subject *domain.Subject) error {
	query := `INSERT INTO subjects (id, user_id, name) VALUES ($1, $2, $3)`
	
	_, err := r.db.ExecContext(ctx, query, subject.ID, subject.UserID, subject.Name)
	return err
}

func (r *subjectRepository) GetByUserID(ctx context.Context, userID string) ([]domain.Subject, error) {
	query := `SELECT id, user_id, name FROM subjects WHERE user_id = $1 ORDER BY name ASC`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subjects []domain.Subject
	for rows.Next() {
		var sub domain.Subject
		err := rows.Scan(&sub.ID, &sub.UserID, &sub.Name)
		if err != nil {
			return nil, err
		}
		subjects = append(subjects, sub)
	}

	return subjects, nil
}