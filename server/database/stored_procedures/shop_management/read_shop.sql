CREATE OR REPLACE FUNCTION read_all_shops() RETURNS SETOF Shops LANGUAGE plpgsql AS 
$$ BEGIN 
    RETURN QUERY SELECT * FROM Shops; 
END; $$;

CREATE OR REPLACE FUNCTION read_shop_by_column(
    IN key_column_name_1 TEXT,
    IN key_column_value_1 TEXT
) RETURNS SETOF Shops LANGUAGE plpgsql AS 
$$ BEGIN
    IF (key_column_name_1 = 'id' OR key_column_name_1 = 'owner_id') THEN
        RETURN QUERY EXECUTE format('SELECT * FROM Shops WHERE %I = $1::INTEGER', key_column_name_1) USING key_column_value_1;
    END IF;
END; $$;