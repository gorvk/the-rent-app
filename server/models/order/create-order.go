package models

import (
	customTypes "github.com/gorvk/rent-app/api-services/common/types"
	"github.com/gorvk/rent-app/api-services/initializers"
)

func CreateOrder(order customTypes.PLACE_ORDER_INPUT) error {
	db := initializers.GetDBInstance()

	stmt, err := db.Prepare(`
		INSERT INTO Orders (
			from_map_location,
			to_map_location,
			last_stop_map_location,
			order_status,
			payment_status,
			product_id,
			buyer_id,
			shop_id,
			quantity
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);
	`)
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(
		order.FromMapLocation,
		order.ToMapLocation,
		order.LastStopMapLocation,
		order.OrderStatus,
		order.PaymentStatus,
		order.ProductId,
		order.BuyerId,
		order.ShopId,
		order.Quantity,
	)
	if err != nil {
		return err
	}

	return nil
}
