CREATE EXTENSION pg_trgm;
CREATE OR REPLACE FUNCTION search_in_products(IN search_term TEXT) RETURNS TABLE (
        shop_id INTEGER,
        product_name VARCHAR,
        product_type VARCHAR,
        product_condition VARCHAR,
        price MONEY,
        product_description VARCHAR,
        id INTEGER,
        shop_name VARCHAR,
        city VARCHAR,
        country VARCHAR,
        document TSVECTOR,
        ss_query TSQUERY,
        rank_product_name REAL,
        similarity REAL
    ) LANGUAGE plpgsql AS $$ BEGIN RETURN QUERY
select *
from search_products_view,
    to_tsvector(
        'english',
        search_products_view.product_description || search_products_view.product_name || search_products_view.product_type || search_products_view.city
    ) AS document,
    to_tsquery('english', search_term) AS s_query,
    ts_rank(
        to_tsvector(
            'english',
            search_products_view.product_description || search_products_view.product_name || search_products_view.product_type || search_products_view.city
        ),
        s_query
    ) AS rank_product_name,
    SIMILARITY(
        search_term,
        search_products_view.product_description || search_products_view.product_name || search_products_view.product_type || search_products_view.city
    ) AS similarity
ORDER BY rank_product_name DESC,
    similarity DESC;
END;
$$;