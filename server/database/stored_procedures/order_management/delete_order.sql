CREATE OR REPLACE PROCEDURE delete_order(
        IN key_column_name TEXT,
        IN key_value TEXT
    ) LANGUAGE plpgsql AS 
$$ BEGIN
    IF key_column_name = 'id' THEN
        EXECUTE format('DELETE FROM Orders WHERE %I = $1::TEXT', key_column_name) USING key_value;
    END IF;
END; $$;