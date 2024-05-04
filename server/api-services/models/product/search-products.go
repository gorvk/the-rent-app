package models

import (
	"database/sql"
	"fmt"

	"github.com/gorvk/rent-app/server/api-services/initializers"
)

func SearchProducts(searchTerm string) (*sql.Rows, error) {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil, nil
	}
	query := fmt.Sprintf("SELECT * from search_in_products('%v')", searchTerm)
	rows, err := db.Query(query)
	return rows, err
}
