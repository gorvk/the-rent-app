CREATE OR REPLACE PROCEDURE update_shop(
        _owner_id INTEGER,
        _shop_name VARCHAR,
        _email VARCHAR,
        _phone_number VARCHAR,
        _map_location VARCHAR,
        _shop_type VARCHAR,
        _shop_description VARCHAR
    ) LANGUAGE 'plpgsql' AS $$ BEGIN EXECUTE format(
        'UPDATE Shops SET %I = $1, %I = $2, %I = $3, %I = $4, %I = $5, %I = $6, %I = $7 WHERE %I = $7', 
        'owner_id', 'shop_name', 'phone_number', 'map_location', 'shop_type', 'shop_description', 'email', 'email'
    ) USING _owner_id, _shop_name, _phone_number, _map_location, _shop_type, _shop_description, _email;
END;
$$;