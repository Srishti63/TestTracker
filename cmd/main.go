package main

import (
	"log"
	"test_tracker_backend/bootstrap"
)

func main() {
	dsn := "postgres://postgres:password@localhost:5432/test_tracker?sslmode=disable"

	db := bootstrap.NewDatabaseConnection(dsn)
	defer db.Close()

	log.Println("Database master key is ready!")
}