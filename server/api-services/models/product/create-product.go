package models

import (
	"database/sql"

	customTypes "github.com/gorvk/rent-app/server/api-services/common/types"
	"github.com/gorvk/rent-app/server/api-services/initializers"
)

func CreateProduct(product customTypes.CREATE_PRODUCT_INPUT, shopId int) (sql.Result, error) {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil, nil
	}

	stmt, err := db.Prepare("CALL create_product($1, $2, $3, $4, $5, $6, $7, $8)")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	rows, err := stmt.Exec(
		product.ProductName,
		shopId,
		product.ProductType,
		product.ProductCondition,
		product.Price,
		product.OriginalPurchasedDate,
		product.OriginalPurchaisingRecieptNo,
		product.ProductDescription,
	)

	return rows, err
}
