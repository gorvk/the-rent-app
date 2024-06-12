package models

import (
	"database/sql"

	customTypes "github.com/gorvk/rent-app/api-services/common/types"
	"github.com/gorvk/rent-app/api-services/initializers"
)

func CreateProduct(product customTypes.CREATE_PRODUCT_INPUT, shopId int) (sql.Result, error) {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil, nil
	}

	stmt, err := db.Prepare(`
		INSERT INTO Products (
			product_name, 
			shop_id, 
			product_type, 
			product_condition, 
			price, 
			original_purchased_date, 
			original_purchaising_reciept_no, 
			product_description
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);
	`)

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

	if err != nil {
		return nil, err
	}

	stmt, err = db.Prepare(`REFRESH MATERIALIZED VIEW search_products_view;`)

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec()

	if err != nil {
		return nil, err
	}

	return rows, err
}
