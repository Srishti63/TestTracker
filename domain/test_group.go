package domain

import "context"

type TestGroup struct{
	TestGroupId string `json:"test_group_id" db:"test_group_id"`
	UserID        string `json:"user_id" db:"user_id"`
	TestGroupName string `json:"test_group_name" db:"test_group_name"`
}

type TestGroupRepository interface{
	Create(ctx context.Context , group *TestGroup) error
	GetByUserId (ctx context.Context,UserID string)([]TestGroup, error)
}

type TestGroupUsecase interface {
	CreateGroup(ctx context.Context, group *TestGroup) error
	FetchUserGroups(ctx context.Context, userID string) ([]TestGroup, error)
}