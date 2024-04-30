package models

import (
	"database/sql"

	"github.com/gorvk/rent-app/server/api-services/initializers"
)

func GetAllUsers() (*sql.Rows, error) {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil, nil
	}
	query := "SELECT * from read_all_users()"
	rows, err := db.Query(query)
	return rows, err
}
