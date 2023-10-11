CREATE TABLE Users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(200) NOT NULL,
    address VARCHAR(200),
    phone VARCHAR(20),
    status VARCHAR(20),
    role VARCHAR(20),
    amount FLOAT,
    registered_at TIMESTAMP
);

CREATE TABLE User_verifications (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    verification_code VARCHAR(30),
    FOREIGN KEY (user_id) REFERENCES Users(id)
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    description VARCHAR(200),
    price FLOAT,
    stock INT
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    order_date TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE order_details (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT,
    total_price FLOAT,
    FOREIGN KEY (order_id) REFERENCES orders(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);


-- contoh record data
INSERT INTO products (name, description, price, stock) VALUES
    ('Ikan Salmon', 'Salmon Norwegia', 750000, 100),
    ('Udang Windu', 'Udang windu besar-besaran dari Indonesia', 450000, 150),
    ('Ikan Tuna', 'Tuna sirip kuning yang segar', 850000, 75),
    ('Kerang Segar', 'Kerang segar dari pantai lokal', 120000, 200),
    ('Kepiting Batik', 'Kepiting batik premium dari Indonesia', 500000, 50);