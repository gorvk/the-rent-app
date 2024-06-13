package models

import (
	"database/sql"
	"fmt"

	"github.com/gorvk/rent-app/api-services/initializers"
)

func SearchProducts(searchTerm string) (*sql.Rows, error) {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil, nil
	}
	query := fmt.Sprintf(`
		select * from search_products_view,
		to_tsvector(
			'english',
			search_products_view.product_description || search_products_view.product_name || search_products_view.product_type || search_products_view.city
		) AS document,
		to_tsquery('english', '%[1]v') AS s_query,
		ts_rank(
			to_tsvector(
				'english',
				search_products_view.product_description || search_products_view.product_name || search_products_view.product_type || search_products_view.city
			)
		,s_query ) AS rank_product_name,
		SIMILARITY(
			'%[1]v',
			search_products_view.product_description || search_products_view.product_name || search_products_view.product_type || search_products_view.city
		) AS similarity
		ORDER BY rank_product_name DESC, similarity DESC;
	`, searchTerm)
	rows, err := db.Query(query)
	return rows, err
}
