package models

import (
	"database/sql"

	customTypes "github.com/gorvk/rent-app/server/api-services/common/types"
	"github.com/gorvk/rent-app/server/api-services/initializers"
)

func UpdateUser(user customTypes.UPDATE_USER_INPUT) (sql.Result, error) {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil, nil
	}

	stmt, err := db.Prepare("CALL update_user($1, $2, $3, $4, $5, $6::BOOLEAN)")

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
