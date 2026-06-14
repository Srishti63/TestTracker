package main

import (
	"log"
	"os"
	"test_tracker_backend/bootstrap"
	"test_tracker_backend/internal/mailutil"
	"test_tracker_backend/repository"
	"test_tracker_backend/route"
	"test_tracker_backend/usecase"
	"time"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {

	err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, falling back to system env variables")
    }

    mailCfg := mailutil.EmailConfig{
        SMTPHost: os.Getenv("SMTP_HOST"),
        SMTPPort: os.Getenv("SMTP_PORT"),
        Sender:   os.Getenv("SMTP_SENDER"),
        Password: os.Getenv("SMTP_PASSWORD"),
    }

	dsn := "postgres://postgres:password@localhost:5432/test_tracker?sslmode=disable"

	db := bootstrap.NewDatabaseConnection(dsn)
	defer db.Close()
	log.Println("Database master key is ready!")

	timeout := 2 * time.Second
	jwtSecret := "your_super_secret_jwt_signing_key"
	jwtExpiryHours := 24

	userRepo := repository.NewUserRepository(db)
	subjectRepo := repository.NewSubjectRepository(db)
	testGroupRepo := repository.NewTestGroupRepository(db)
	presetRepo := repository.NewSubjectPresetRepository(db)
	testRepo := repository.NewTestRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepo,timeout, jwtSecret, jwtExpiryHours, &mailCfg)
	subjectUsecase := usecase.NewSubjectUsecase(subjectRepo,timeout) 
	testGroupUsecase := usecase.NewTestGroupUsecase(testGroupRepo,timeout)
	presetUsecase := usecase.NewSubjectPresetUsecase(presetRepo,timeout)
	testUsecase := usecase.NewTestUsecase(testRepo,timeout)

	r := gin.Default()

	// calling the distinct routes distinctly 
	route.SetupUserRoutes(r, userUsecase)
	route.SetupSubjectRoutes(r, subjectUsecase)
	route.SetupTestGroupRoutes(r, testGroupUsecase)
	route.SetupSubjectPresetRoutes(r, presetUsecase)
	route.SetupTestRoutes(r, testUsecase)

	log.Println("Backend engine running smoothly on port :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Critical system failure launching server: %v", err)
	}
}
