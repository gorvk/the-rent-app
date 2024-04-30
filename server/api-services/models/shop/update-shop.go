package models

import (
	"database/sql"

	customTypes "github.com/gorvk/rent-app/server/api-services/common/types"
	"github.com/gorvk/rent-app/server/api-services/initializers"
)

func UpdateShop(shop customTypes.UPDATE_SHOP_INPUT, ownerId int) (sql.Result, error) {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil, nil
	}

	stmt, err := db.Prepare("CALL update_shop($1, $2, $3, $4, $5, $6, $7)")

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
