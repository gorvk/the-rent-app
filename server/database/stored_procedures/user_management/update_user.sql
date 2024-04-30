CREATE OR REPLACE PROCEDURE update_user(
        _email VARCHAR,
        _first_name VARCHAR,
        _last_name VARCHAR,
        _phone_number VARCHAR,
        _user_address VARCHAR,
        _is_shop_enabled BOOLEAN
    ) LANGUAGE 'plpgsql' AS $$ BEGIN EXECUTE format(
        'UPDATE Users SET %I = $1, %I = $2, %I = $3, %I = $4, %I = $5, %I = $6 WHERE %I = $1', 
        'email', 'first_name', 'last_name', 'phone_number', 'user_address', 'is_shop_enabled', 'email'
    ) USING _email, _first_name, _last_name, _phone_number, _user_address, _is_shop_enabled;
END;
$$;