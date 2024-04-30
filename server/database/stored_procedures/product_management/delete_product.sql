CREATE OR REPLACE PROCEDURE delete_product (
        IN key_column_name TEXT,
        IN key_value TEXT
    ) LANGUAGE plpgsql AS 
$$ BEGIN
    IF key_column_name = 'id' THEN
        EXECUTE format('DELETE FROM Products WHERE %I = $1::TEXT', key_column_name) USING key_value;
    END IF;
END; $$;