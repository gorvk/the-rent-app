package models

import "github.com/gorvk/rent-app/api-services/initializers"

func RefreshProductViews() error {
	db := initializers.GetDBInstance()

	stmt, err := db.Prepare(`REFRESH MATERIALIZED VIEW product_detail_view;`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	stmt, err = db.Prepare(`REFRESH MATERIALIZED VIEW search_products_view;`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	return nil
}
