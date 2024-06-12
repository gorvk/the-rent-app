package models

import (
	"database/sql"
	"fmt"

	"github.com/gorvk/rent-app/api-services/initializers"
)

func DeleteUser(email string) (*sql.Rows, error) {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil, nil
	}
	query := fmt.Sprintf("DELETE FROM Users WHERE email = %v", email)
	rows, err := db.Query(query)
	return rows, err
}
