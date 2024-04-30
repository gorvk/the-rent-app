package models

import (
	"database/sql"

	"github.com/gorvk/rent-app/server/api-services/initializers"
)

func GetAllProducts() (*sql.Rows, error) {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil, nil
	}
	query := "SELECT * from read_all_products()"
	rows, err := db.Query(query)
	return rows, err
}
