package models

import (
	"database/sql"
	"fmt"

	"github.com/gorvk/rent-app/api-services/initializers"
)

func DeleteProduct(productId int, shopId int) (*sql.Rows, error) {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil, nil
	}
	query := fmt.Sprintf("CALL delete_product('%v', '%v', '%v', '%v')", "id", productId, "shop_id", shopId)
	rows, err := db.Query(query)
	return rows, err
}
