package bootstrap

import(
	"database/sql"
	"log"
	"time"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewDatabaseConnection(envDataSource string) *sql.DB{

	db, err := sql.Open("pgx", envDataSource)
	if err != nil {
		log.Fatalf("failed to open the database connection", err);
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil{
		log.Fatalf("Database connection failed (.Ping)  : %v",err);
	}

	log.Println("Database connection pool successfully established !!");
	return db
}