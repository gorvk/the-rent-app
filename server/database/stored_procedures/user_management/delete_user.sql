CREATE OR REPLACE PROCEDURE delete_user(
        IN key_column_name TEXT,
        IN key_value TEXT
    ) LANGUAGE plpgsql AS 
$$ BEGIN
    IF key_column_name = 'email' THEN
        EXECUTE format('DELETE FROM Users WHERE %I = $1::TEXT', key_column_name) USING key_value;
    END IF;
END; $$;