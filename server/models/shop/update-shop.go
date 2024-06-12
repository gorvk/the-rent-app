package models

import (
	"database/sql"

	customTypes "github.com/gorvk/rent-app/api-services/common/types"
	"github.com/gorvk/rent-app/api-services/initializers"
)

func UpdateShop(shop customTypes.UPDATE_SHOP_INPUT, ownerId int) (sql.Result, error) {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil, nil
	}

	stmt, err := db.Prepare(`
		UPDATE Shops SET 
			owner_id = $1, 
			shop_name = $2, 
			phone_number = $3, 
			map_location = $4, 
			shop_type = $5, 
			shop_description = $6, 
			email = $7 
		WHERE email = $7
	`)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	rows, err := stmt.Exec(
		ownerId,
		shop.ShopName,
		shop.Email,
		shop.PhoneNumber,
		shop.MapLocation,
		shop.ShopType,
		shop.ShopDescription,
	)

	return rows, err
}
