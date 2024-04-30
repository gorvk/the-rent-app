CREATE OR REPLACE PROCEDURE update_product (
    _id INTEGER,
    _product_name VARCHAR,
    _shop_id INTEGER,
    _product_type VARCHAR,
    _product_condition VARCHAR,
    _price MONEY,
    _original_purchased_date DATE,
    _original_purchaising_reciept_no VARCHAR,
    _product_description VARCHAR
) LANGUAGE 'plpgsql' AS $$ BEGIN EXECUTE format(
        'UPDATE Products SET %I = $1, %I = $2, %I = $3, %I = $4, %I = $5, %I = $6, %I = $7, %I = $8 WHERE %I = $9', 
        'product_name', 'shop_id', 'product_type', 'product_condition', 'price', 'original_purchased_date', 'original_purchaising_reciept_no', 'product_description', 'id'
    ) USING _product_name, _shop_id, _product_type, _product_condition, _price, _original_purchased_date, _original_purchaising_reciept_no, _product_description, _id;
    END;
$$;