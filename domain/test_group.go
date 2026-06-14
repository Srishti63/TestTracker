package domain

import "context"

type TestGroup struct {
	TestGroupID   string `json:"test_group_id" db:"id"`
	UserID        string `json:"user_id" db:"user_id"`
	TestGroupName string `json:"test_group_name" db:"test_group_name"`
}

type CreateTestGroupParam struct {
	TestGroupName string
	UserID        string
}

type TestGroupRepository interface {
	Create(ctx context.Context, group *TestGroup) error
	GetByByUserId(ctx context.Context, userID string) ([]TestGroup, error)
}

type TestGroupUsecase interface {
	CreateGroup(ctx context.Context, req *CreateTestGroupParam) (*TestGroup,error)
	FetchUserGroups(ctx context.Context, userID string) ([]TestGroup, error)
}