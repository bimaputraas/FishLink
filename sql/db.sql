CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name TEXT,
    description TEXT,
    price FLOAT,
    stock INT
);

CREATE TABLE Users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(200) NOT NULL,
    address VARCHAR(100),
    phone VARCHAR(20),
    status VARCHAR(20),
    registered_at TIMESTAMP
);

CREATE TABLE User_verifications (
    id SERIAL PRIMARY KEY,
    user_id INT,
    verification_code VARCHAR(30),
    FOREIGN KEY (user_id) REFERENCES Users(id)
);