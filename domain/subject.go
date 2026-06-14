package domain

import (
	"context"
)

type Subject struct {
	ID     string `json:"id" db:"id"`
	UserID string `json:"user_id" db:"user_id"` // Keeps catalog private
	Name   string `json:"name" db:"name"`
}

type CreateSubjectParam struct{
	Name string
	UserID string
}

type SubjectRepository interface {
	Create(ctx context.Context, subject *Subject) error
	GetByUserID(ctx context.Context, userID string) ([]Subject, error)
}

type SubjectUsecase interface {
    CreateSubject(ctx context.Context, req *CreateSubjectParam) (*Subject, error) 
    FetchUserSubjects(ctx context.Context, userID string) ([]Subject, error)
}