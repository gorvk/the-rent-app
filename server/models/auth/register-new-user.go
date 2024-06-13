package models

import (
	"database/sql"

	customTypes "github.com/gorvk/rent-app/api-services/common/types"
	"github.com/gorvk/rent-app/api-services/initializers"
)

func RegisterNewUser(user customTypes.User, saltHashedPassword []byte) (sql.Result, error) {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil, nil
	}

	stmt, err := db.Prepare(`
		INSERT INTO Users (
			email, 
			first_name, 
			last_name, 
			phone_number, 
			user_address, 
			account_password
		) VALUES ($1, $2, $3, $4, $5, $6::BYTEA);
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
		saltHashedPassword,
	)

	return rows, err
}
