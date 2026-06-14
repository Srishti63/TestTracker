package usecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"test_tracker_backend/domain"
	"time"
)

type subjectUsecase struct {
	subjectRepo    domain.SubjectRepository
	contextTimeout time.Duration
}

func NewSubjectUsecase(sr domain.SubjectRepository, timeout time.Duration) domain.SubjectUsecase {
	return &subjectUsecase{
		subjectRepo:    sr,
		contextTimeout: timeout,
	}
}

func (u *subjectUsecase) CreateSubject(ctx context.Context, req *domain.CreateSubjectParam) (*domain.Subject, error) {
	c, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	if req.Name == "" || req.UserID == "" {
		return nil, errors.New("subject name and user credentials cannot be blank")
	}

	subject := &domain.Subject{
		ID:     uuid.New().String(),
		UserID: req.UserID,
		Name:   req.Name,
	}

	err := u.subjectRepo.Create(c, subject)
	if err != nil {
		return nil, err 
	}
	return subject, nil
}

func (u *subjectUsecase) FetchUserSubjects(ctx context.Context, userID string) ([]domain.Subject, error) {
	c, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	if userID == "" {
		return nil, errors.New("user identification context tracking parameter missing")
	}

	return u.subjectRepo.GetByUserID(c, userID)
}
