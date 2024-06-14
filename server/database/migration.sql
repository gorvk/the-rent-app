CREATE TABLE IF NOT EXISTS Users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(15) NOT NULL,
    last_name VARCHAR(15) NOT NULL,
    email VARCHAR(20) UNIQUE NOT NULL,
    phone_number VARCHAR(20) UNIQUE NOT NULL,
    user_address VARCHAR NOT NULL,
    is_shop_enabled BOOLEAN DEFAULT FALSE,
    account_password BYTEA NOT NULL
);

CREATE TABLE IF NOT EXISTS Shops (
    id SERIAL PRIMARY KEY,
    owner_id INTEGER NOT NULL,
    shop_name VARCHAR(100) NOT NULL,
    email VARCHAR UNIQUE NOT NULL,
    phone_number VARCHAR UNIQUE,
    map_location VARCHAR NOT NULL,
    shop_type VARCHAR(100) NOT NULL,
    shop_description VARCHAR(500) NOT NULL,
    city VARCHAR NOT NULL,
    country VARCHAR NOT NULL,
    FOREIGN KEY (owner_id) REFERENCES Users(id)
);

CREATE TABLE IF NOT EXISTS Products(
    id SERIAL PRIMARY KEY,
    product_name VARCHAR(100) NOT NULL,
    shop_id SERIAL NOT NULL,
    product_type VARCHAR(100),
    product_condition VARCHAR(20),
    price MONEY,
    original_purchased_date DATE,
    original_purchaising_reciept_no VARCHAR,
    product_description VARCHAR(1000),
    FOREIGN KEY (shop_id) REFERENCES Shops(id)
);

CREATE TABLE IF NOT EXISTS Orders (
    id SERIAL PRIMARY KEY,
    from_map_location VARCHAR NOT NULL,
    to_map_location VARCHAR NOT NULL,
    last_stop_map_location VARCHAR NOT NULL,
    order_status VARCHAR(10) NOT NULL,
    payment_status VARCHAR NOT NULL,
    product_id SERIAL NOT NULL,
    buyer_id SERIAL NOT NULL,
    shop_id SERIAL NOT NULL,
    FOREIGN KEY (product_id) REFERENCES Products(id),
    FOREIGN KEY (buyer_id) REFERENCES Users(id),
    FOREIGN KEY (shop_id) REFERENCES Shops(id)
);

CREATE EXTENSION pg_trgm;

CREATE MATERIALIZED VIEW IF NOT EXISTS public.search_products_view TABLESPACE pg_default AS
    SELECT 
        product_name,
        products.product_type,
        products.product_condition,
        products.price,
        products.product_description,
        shops.id shop_id,
        shops.shop_name,
        shops.city,
        shops.country
    FROM products
        LEFT JOIN shops ON products.shop_id = shops.id;

CREATE INDEX IF NOT EXISTS search_products_view_product_name ON search_products_view USING GIN (
    to_tsvector(
        'english',
        search_products_view.product_name || search_products_view.product_description || search_products_view.product_type || search_products_view.city
    )
);

