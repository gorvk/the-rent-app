package models

import (
	"database/sql"
	"fmt"

	"github.com/gorvk/rent-app/server/api-services/initializers"
)

func GetCurrentShopProduct(shopId int) (*sql.Rows, error) {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil, nil
	}
	query := fmt.Sprintf("select * from read_product_by_column('%v', '%v')", "shop_id", shopId)
	rows, err := db.Query(query)
	return rows, err
}
