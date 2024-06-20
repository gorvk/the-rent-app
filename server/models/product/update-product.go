package models

import (
	customTypes "github.com/gorvk/rent-app/api-services/common/types"
	"github.com/gorvk/rent-app/api-services/initializers"
)

func UpdateProduct(product customTypes.Product) error {
	db := initializers.GetDBInstance()

	stmt, err := db.Prepare(`
		UPDATE Products SET 
			product_name = $1, 
			shop_id = $2, 
			product_type = $3, 
			product_condition = $4, 
			price = $5, 
			original_purchased_date = $6, 
			original_purchaising_reciept_no = $7, 
			product_description = $8,
			quantity = $9
		WHERE id = $10;
	`)

	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(
		product.ProductName,
		product.ShopId,
		product.ProductType,
		product.ProductCondition,
		product.Price,
		product.OriginalPurchasedDate,
		product.OriginalPurchaisingRecieptNo,
		product.ProductDescription,
		product.Quantity,
		product.Id,
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
