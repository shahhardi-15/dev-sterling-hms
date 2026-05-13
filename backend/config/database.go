package config

import (
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func ConnectDB() *sqlx.DB {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set in .env")
	}

	db, err := sqlx.Connect("pgx", dsn+"?default_query_exec_mode=simple_protocol")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	log.Println("Database connected successfully!")
	DB = db
	return db
}
