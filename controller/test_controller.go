package controller

import (
	"database/sql"
	"net/http"
	"test_tracker_backend/domain"
)

type testController struct{
	db *sql.DB
}

func NewTestController (test domain.TestUsecase) *testController{
	
} 