package models

import (
	"database/sql"
	"fmt"

	"github.com/gorvk/rent-app/server/api-services/initializers"
)

func GetShopByOwnerId(shopId int) (*sql.Rows, error) {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil, nil
	}
	query := fmt.Sprintf("SELECT * FROM read_shop_by_column('%v', '%v')", "owner_id", shopId)
	rows, err := db.Query(query)
	return rows, err
}
