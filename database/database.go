package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/godror/godror"
	"github.com/joho/godotenv"
	_ "github.com/sijms/go-ora/v2"
)

func Initialize() *sql.DB {
	db, err := sql.Open("oracle", generateConnectionString())
	if err != nil {
		panic(fmt.Errorf("error in while opening the connection: %w", err))
	}

	err = db.Ping()
	if err != nil {
		panic(fmt.Errorf("error pinging db: %w", err))
	}

	fmt.Println("database got successfully connected")
	return db
}

func generateConnectionString() string {
	err := godotenv.Load("../.envrc")
	if err != nil {
		panic(fmt.Errorf("Error loading .envrc file", err))
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbServiceName := os.Getenv("DB_SERVICE_NAME")

	return fmt.Sprintf("oracle://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbServiceName)
}
