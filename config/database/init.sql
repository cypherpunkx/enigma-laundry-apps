CREATE DATABASE enigma_laundry_v3;

CREATE TABLE employees (
    id VARCHAR(100) PRIMARY KEY NOT NULL,
    name VARCHAR(100),
    phone_number VARCHAR(20) UNIQUE,
    address TEXT
);

CREATE TABLE products (
    id VARCHAR(100) PRIMARY KEY NOT NULL,
    name VARCHAR(100) NOT NULL,
    price BIGINT,
    uom VARCHAR(50)
);

CREATE TABLE customers (
    id VARCHAR(100) PRIMARY KEY NOT NULL,
    name VARCHAR(100),
    phone_number VARCHAR(20) UNIQUE,
    address TEXT
);


CREATE TABLE bills (
    id VARCHAR(100) PRIMARY KEY NOT NULL,
    bill_date DATE,
    entry_date DATE,
    finish_date DATE,
    employee_id VARCHAR(100),
    customer_id VARCHAR(100),
    FOREIGN KEY(employee_id) REFERENCES employees(id),
    FOREIGN KEY(customer_id) REFERENCES customers(id)
);

CREATE TABLE bill_details (
    id VARCHAR(100) PRIMARY KEY NOT NULL,
    bill_id VARCHAR(100),
    product_id VARCHAR(100),
    product_price BIGINT,
    qty INT,
    FOREIGN KEY(bill_id) REFERENCES bills(id),
    FOREIGN KEY(product_id) REFERENCES products(id)
);

CREATE TABLE users (
        id VARCHAR(100) PRIMARY KEY NOT NULL,
        username VARCHAR(50) NOT NULL UNIQUE,
        password VARCHAR(200) NOT NULL,
        is_active BOOLEAN DEFAULT true
);