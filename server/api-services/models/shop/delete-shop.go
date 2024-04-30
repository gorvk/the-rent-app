package models

import (
	"database/sql"
	"fmt"

	"github.com/gorvk/rent-app/server/api-services/initializers"
)

func DeleteShop(shopEmail string, ownerId int) (*sql.Rows, error) {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil, nil
	}
	query := fmt.Sprintf("CALL delete_shop('%v', '%v', '%v', '%v')", "email", shopEmail, "owner_id", ownerId)
	rows, err := db.Query(query)
	return rows, err
}
