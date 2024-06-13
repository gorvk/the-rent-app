package initializers

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectDB() error {
	fmt.Println("Connecting to DB...")

	var err error
	DB_CONNECTION_STRING := os.Getenv("DB_CONNECTION_STRING")

	// opening the connection with Postgres DB
	db, err = sql.Open("postgres", DB_CONNECTION_STRING)
	if err != nil {
		return err
	}

	// checking the ping to DB
	err = db.Ping()
	if err != nil {
		return err
	}

	fmt.Println("DB Connected Successfully !")
	return nil
}

func DbMigration() error {
	fmt.Println("DB Migration started...")

	file, err := os.ReadFile("database/migration.sql")
	if err != nil {
		return err
	}

	_, err = db.Exec(string(file))
	if err != nil {
		return err
	}

	fmt.Println("DB Migration Completed !")
	return nil
}

func GetDBInstance() *sql.DB {
	return db
}
