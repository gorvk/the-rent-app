package models

import (
	"database/sql"
	"fmt"

	"github.com/gorvk/rent-app/api-services/initializers"
)

func GetShopByOwnerId(ownerId int) (*sql.Rows, error) {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil, nil
	}
	query := fmt.Sprintf("SELECT * FROM Shops WHERE owner_id = %v", ownerId)
	rows, err := db.Query(query)
	return rows, err
}
