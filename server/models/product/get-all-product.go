package models

import (
	"database/sql"

	"github.com/gorvk/rent-app/api-services/initializers"
)

func GetAllProducts() (*sql.Rows, error) {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil, nil
	}
	query := "SELECT * FROM search_products_view"
	rows, err := db.Query(query)
	return rows, err
}
