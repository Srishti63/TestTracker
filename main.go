package main

import (
	"log"
	"time"
	"test_tracker_backend/bootstrap"
	"test_tracker_backend/repository" 
	"test_tracker_backend/route"

	"github.com/gin-gonic/gin"
)

func main() {
	dsn := "postgres://postgres:password@localhost:5432/test_tracker?sslmode=disable"

	db := bootstrap.NewDatabaseConnection(dsn)
	defer db.Close()
	log.Println("Database master key is ready!")

	userRepo := repository.NewUserRepository(db)

	r := gin.Default()

	timeout := 2 * time.Second
	jwtSecret := "your_super_secret_jwt_signing_key"
	jwtExpiryHours := 24

	route.Setup(r, userRepo, timeout, jwtSecret, jwtExpiryHours)

	log.Println("Backend engine running smoothly on port :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Critical system failure launching server: %v", err)
	}
}