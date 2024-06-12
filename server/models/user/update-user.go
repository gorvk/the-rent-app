package models

import (
	"database/sql"

	customTypes "github.com/gorvk/rent-app/api-services/common/types"
	"github.com/gorvk/rent-app/api-services/initializers"
)

func UpdateUser(user customTypes.UPDATE_USER_INPUT) (sql.Result, error) {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil, nil
	}

	// UPDATE Users SET %I = $1, %I = $2, %I = $3, %I = $4, %I = $5, %I = $6 WHERE %I = $1
	stmt, err := db.Prepare(`
		UPDATE Users SET 
			email=$1,
			first_name=$2,
			last_name=$3,
			phone_number=$4,
			user_address=$5,
			is_shop_enabled=$6
		WHERE email=$1
	`)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	rows, err := stmt.Exec(
		user.Email,
		user.FirstName,
		user.LastName,
		user.PhoneNumber,
		user.UserAddress,
		user.IsShopEnabled,
	)

	return rows, err
}
