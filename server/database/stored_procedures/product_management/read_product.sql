CREATE OR REPLACE FUNCTION read_all_products() RETURNS TABLE (
    id INTEGER,
    product_name VARCHAR,
    shop_id INTEGER,
    product_type VARCHAR,
    product_condition VARCHAR,
    price MONEY,
    original_purchased_date DATE,
    original_purchaising_reciept_no VARCHAR,
    product_description VARCHAR
) LANGUAGE plpgsql AS 
$$ BEGIN 
    RETURN QUERY SELECT * FROM Products; 
END; $$;

CREATE OR REPLACE FUNCTION read_product_by_column(
    IN key_column_name TEXT,
    IN key_value TEXT
) RETURNS TABLE (
    id INTEGER,
    product_name VARCHAR,
    shop_id INTEGER,
    product_type VARCHAR,
    product_condition VARCHAR,
    price MONEY,
    original_purchased_date DATE,
    original_purchaising_reciept_no VARCHAR,
    product_description VARCHAR
    ) LANGUAGE plpgsql AS 
$$ BEGIN    
    IF key_column_name = 'id' THEN
        RETURN QUERY EXECUTE format('SELECT * FROM Products WHERE %I = $1::INTEGER', key_column_name) USING key_value;
    ELSE
        RETURN QUERY EXECUTE format('SELECT * FROM Products WHERE %I = $1', key_column_name) USING key_value;
    END IF;
END; $$;