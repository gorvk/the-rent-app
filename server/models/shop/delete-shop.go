package models

import (
	"database/sql"
	"fmt"

	"github.com/gorvk/rent-app/api-services/initializers"
)

func DeleteShop(shopEmail string, ownerId int) (*sql.Rows, error) {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil, nil
	}

	query := fmt.Sprintf(`DELETE FROM Shops WHERE email = %v AND owner_id = %v`, shopEmail, ownerId)
	rows, err := db.Query(query)
	return rows, err
}
