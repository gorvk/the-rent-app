CREATE MATERIALIZED VIEW IF NOT EXISTS public.search_products_view TABLESPACE pg_default AS
SELECT 
    product_name,
    products.product_type,
    products.product_condition,
    products.price,
    products.product_description,
	shops.id shop_id,
    shops.shop_name,
    shops.city,
    shops.country
FROM products
    LEFT JOIN shops ON products.shop_id = shops.id;
CREATE INDEX search_products_view_product_name ON search_products_view USING GIN (
    to_tsvector(
        'english',
        search_products_view.product_name || search_products_view.product_description || search_products_view.product_type || search_products_view.city
    )
);