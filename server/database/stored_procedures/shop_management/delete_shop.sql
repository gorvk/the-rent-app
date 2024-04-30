CREATE OR REPLACE PROCEDURE delete_shop(
        IN key_column_name_1 TEXT,
        IN key_column_value_1 TEXT,        
        IN key_column_name_2 TEXT,
        IN key_column_value_2 TEXT
    ) LANGUAGE plpgsql AS 
$$ BEGIN
    IF (key_column_name_1 = 'email') AND (key_column_name_2 = 'owner_id') THEN
        EXECUTE format('DELETE FROM Shops WHERE %I = $1::TEXT AND %I = $2::INTEGER', key_column_name_1, key_column_name_2) USING key_column_value_1, key_column_value_2;
    END IF;
END; $$;