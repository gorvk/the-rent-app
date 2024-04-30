CREATE OR REPLACE PROCEDURE update_order(
    _id INTEGER,
    _from_map_location VARCHAR,
    _to_map_location VARCHAR,
    _last_stop_map_location VARCHAR,
    _order_status VARCHAR,
    _payment_status VARCHAR,
    _product_id INTEGER,
    _buyer_id INTEGER,
    _shop_id INTEGER
) LANGUAGE 'plpgsql' AS $$ BEGIN EXECUTE format(
        'UPDATE Orders SET %I = $1, %I = $2, %I = $3, %I = $4, %I = $5, %I = $6, %I = $7, %I = $8 WHERE %I = $9', 
        'from_map_location', 'to_map_location', 'last_stop_map_location', 'order_status', 'payment_status', 'product_id', 'buyer_id', 'shop_id', 'id'
    ) USING  _from_map_location, _to_map_location, _last_stop_map_location, _order_status, _payment_status, _product_id, _buyer_id, _shop_id, _id;
END;
$$;