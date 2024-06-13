package models

import (
	"database/sql"

	customTypes "github.com/gorvk/rent-app/api-services/common/types"
	"github.com/gorvk/rent-app/api-services/initializers"
)

func UpdateProduct(product customTypes.Product) (sql.Result, error) {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil, nil
	}

	stmt, err := db.Prepare(`
		UPDATE Products SET 
			product_name = $1, 
			shop_id = $2, 
			product_type = $3, 
			product_condition = $4, 
			price = $5, 
			original_purchased_date = $6, 
			original_purchaising_reciept_no = $7, 
			product_description = $8 
		WHERE id = $9;
	`)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	rows, err := stmt.Exec(
		product.ProductName,
		product.ShopId,
		product.ProductType,
		product.ProductCondition,
		product.Price,
		product.OriginalPurchasedDate,
		product.OriginalPurchaisingRecieptNo,
		product.ProductDescription,
		product.Id,
	)

	return rows, err
}
