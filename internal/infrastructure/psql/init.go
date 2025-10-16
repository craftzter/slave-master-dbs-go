package psql

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func InitMaster() (*sql.DB, error) {
	dbUser := os.Getenv("POSTGRES_USER")
	dbName := os.Getenv("POSTGRES_DB")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	dbHost := os.Getenv("DB_PRIMARY_HOST")
	dbPort := os.Getenv("DB_PRIMARY_PORT")
	if dbUser == "" || dbPass == "" || dbName == "" || dbHost == "" || dbPort == "" {
		return nil, fmt.Errorf("missing required database environment variables")
	}
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s host=%s sslmode=disable", dbUser, dbPass, dbName, dbPort, dbHost)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to postgres master %w", err)
	}
	db.SetMaxOpenConns(25)                 // Max connection pool
	db.SetMaxIdleConns(5)                  // Connection yang idle
	db.SetConnMaxLifetime(5 * time.Minute) // Lifetime setiap connection
	log.Println("Postgresql master connected successfully")
	return db, nil
}

func InitSlave() (*sql.DB, error) {
	dbUser := os.Getenv("POSTGRES_USER")
	dbName := os.Getenv("POSTGRES_DB")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	dbHost := os.Getenv("DB_REPLICA_HOST")
	dbPort := os.Getenv("DB_REPLICA_PORT")
	if dbUser == "" || dbPass == "" || dbName == "" || dbHost == "" || dbPort == "" {
		return nil, fmt.Errorf("missing required database environment variables")
	}
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s host=%s sslmode=disable", dbUser, dbPass, dbName, dbPort, dbHost)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to postgres slave %w", err)
	}
	db.SetMaxOpenConns(25)                 // Max connection pool
	db.SetMaxIdleConns(5)                  // Connection yang idle
	db.SetConnMaxLifetime(5 * time.Minute) // Lifetime setiap connection
	log.Println("Postgresql slave connected successfully")
	return db, nil
}
