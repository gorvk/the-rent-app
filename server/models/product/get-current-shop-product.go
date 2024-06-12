package models

import (
	"database/sql"
	"fmt"

	"github.com/gorvk/rent-app/api-services/initializers"
)

func GetCurrentShopProduct(id int) (*sql.Rows, error) {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil, nil
	}
	query := fmt.Sprintf("SELECT * FROM search_products_view WHERE shop_id = %v", id)
	rows, err := db.Query(query)
	return rows, err
}
