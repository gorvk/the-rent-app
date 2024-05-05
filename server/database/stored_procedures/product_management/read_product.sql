CREATE OR REPLACE FUNCTION read_all_products() RETURNS SETOF search_products_view LANGUAGE plpgsql AS 
$$ BEGIN 
    RETURN QUERY SELECT * FROM search_products_view; 
END; $$;

CREATE OR REPLACE FUNCTION read_product_by_column(
    IN key_column_name TEXT,
    IN key_value TEXT
) RETURNS SETOF search_products_view LANGUAGE plpgsql AS 
$$ BEGIN    
    IF key_column_name = 'id' OR key_column_name = 'shop_id' THEN
        RETURN QUERY EXECUTE format('SELECT * FROM search_products_view WHERE %I = $1::INTEGER', key_column_name) USING key_value;
    ELSE
        RETURN QUERY EXECUTE format('SELECT * FROM search_products_view WHERE %I = $1', key_column_name) USING key_value;
    END IF;
END; $$;