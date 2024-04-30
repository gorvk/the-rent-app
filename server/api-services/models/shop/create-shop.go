package models

import (
	"database/sql"

	customTypes "github.com/gorvk/rent-app/server/api-services/common/types"
	"github.com/gorvk/rent-app/server/api-services/initializers"
)

func CreateNewShop(shop customTypes.CREATE_SHOP_INPUT, userId int) (sql.Result, error) {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil, nil
	}

	stmt, err := db.Prepare("CALL create_shop($1, $2, $3, $4, $5, $6, $7)")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	rows, err := stmt.Exec(
		userId,
		shop.ShopName,
		shop.Email,
		shop.PhoneNumber,
		shop.MapLocation,
		shop.ShopType,
		shop.ShopDescription,
	)

	return rows, err
}
