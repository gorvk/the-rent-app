package models

import (
	customTypes "github.com/gorvk/rent-app/api-services/common/types"
	"github.com/gorvk/rent-app/api-services/initializers"
)

func CreateProduct(product customTypes.CREATE_PRODUCT_INPUT, shopId int) error {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil
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
			product_description,
			quantity
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);
	`)
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(
		product.ProductName,
		shopId,
		product.ProductType,
		product.ProductCondition,
		product.Price,
		product.OriginalPurchasedDate,
		product.OriginalPurchaisingRecieptNo,
		product.ProductDescription,
		product.Quantity,
	)
	if err != nil {
		return err
	}

	err = RefreshProductViews()
	if err != nil {
		return err
	}

	return nil
}
