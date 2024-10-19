CREATE TABLE ProductDescription (
    id SERIAL PRIMARY KEY,
    size VARCHAR(50),
    color VARCHAR(50),
    material VARCHAR(100),
    gender VARCHAR(50),
    additional_info TEXT
);

CREATE TABLE ProductCategory (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE ProductImage (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    image VARCHAR(255) NOT NULL
);

CREATE TABLE Customer (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    image_id INTEGER REFERENCES ProductImage(id),
    price DECIMAL(10, 2) NOT NULL,
    description_id INTEGER REFERENCES ProductDescription(id),
    category_id INTEGER REFERENCES ProductCategory(id)
);

CREATE TABLE Orders (
    id SERIAL PRIMARY KEY,
    customer_id INTEGER REFERENCES Customer(id),
    order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    total_amount DECIMAL(10, 2) NOT NULL,
    status VARCHAR(50) NOT NULL
);

CREATE TABLE OrderItems (
    order_id INTEGER REFERENCES Orders(id),
    product_id INTEGER REFERENCES Products(id),
    quantity INTEGER NOT NULL,
    PRIMARY KEY (order_id, product_id)
);

CREATE TABLE Favorites (
    id SERIAL PRIMARY KEY,
    customer_id INTEGER REFERENCES Customer(id),
    product_id INTEGER REFERENCES Products(id)
);

CREATE TABLE Reviews (
    id SERIAL PRIMARY KEY,
    product_id INTEGER REFERENCES Products(id),
    customer_id INTEGER REFERENCES Customer(id),
    rating INTEGER CHECK (rating >= 1 AND rating <= 5),
    comment TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    order_id INTEGER REFERENCES Orders(id)
);

CREATE TABLE Carts (
    id SERIAL PRIMARY KEY,
    customer_id INTEGER REFERENCES Customer(id)
);

CREATE TABLE CartItems (
    cart_id INTEGER REFERENCES Carts(id),
    product_id INTEGER REFERENCES Products(id),
    quantity INTEGER NOT NULL CHECK (quantity > 0),
    added_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (cart_id, product_id)
);