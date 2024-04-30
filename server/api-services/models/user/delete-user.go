package models

import (
	"database/sql"
	"fmt"

	"github.com/gorvk/rent-app/server/api-services/initializers"
)

func DeleteUser(email string) (*sql.Rows, error) {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil, nil
	}
	query := fmt.Sprintf("CALL delete_user('%v', '%v')", "email", email)
	rows, err := db.Query(query)
	return rows, err
}
