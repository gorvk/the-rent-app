DROP TABLE IF EXISTS Shops CASCADE;
CREATE TABLE Shops (
    id SERIAL PRIMARY KEY,
    owner_id INTEGER NOT NULL,
    shop_name VARCHAR(100) NOT NULL,
    email VARCHAR UNIQUE NOT NULL,
    phone_number VARCHAR UNIQUE,
    map_location VARCHAR NOT NULL,
    shop_type VARCHAR(100) NOT NULL,
    shop_description VARCHAR(500) NOT NULL,
    FOREIGN KEY (owner_id) REFERENCES Users(id)
);