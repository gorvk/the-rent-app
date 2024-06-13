package models

import (
	"database/sql"

	customTypes "github.com/gorvk/rent-app/api-services/common/types"
	"github.com/gorvk/rent-app/api-services/initializers"
)

func CreateNewShop(shop customTypes.CREATE_SHOP_INPUT, ownerId int) (sql.Result, error) {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil, nil
	}

	stmt, err := db.Prepare(`
		INSERT INTO Shops (owner_id, shop_name, email, phone_number, map_location, shop_type, shop_description, city, country) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
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
		shop.City,
		shop.Country,
	)

	return rows, err
}
