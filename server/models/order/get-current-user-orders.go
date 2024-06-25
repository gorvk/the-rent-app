package models

import (
	"database/sql"
	"fmt"

	"github.com/gorvk/rent-app/api-services/initializers"
)

func GetCurrentUserOrders(userId int) (*sql.Rows, error) {
	db := initializers.GetDBInstance()

	query := fmt.Sprintf("SELECT * FROM order_detail_view WHERE buyer_id = %v", userId)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
