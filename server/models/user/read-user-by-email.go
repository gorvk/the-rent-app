package models

import (
	"database/sql"
	"fmt"

	"github.com/gorvk/rent-app/api-services/initializers"
)

func GetUserByEmail(email string) (*sql.Rows, error) {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil, nil
	}
	query := fmt.Sprintf("SELECT * FROM Users WHERE email = '%v';", email)
	fmt.Println(query)
	rows, err := db.Query(query)
	return rows, err
}
