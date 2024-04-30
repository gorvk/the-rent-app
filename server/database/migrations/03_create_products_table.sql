DROP TABLE IF EXISTS Products CASCADE;
CREATE TABLE Products(
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