-- Creating Database named psql_assignment
CREATE DATABASE psql_assignment;

-- Create 3 tables
-- Table customers
CREATE TABLE IF NOT EXISTS customers (
  id SERIAL PRIMARY KEY,
  fullname char(50) NOT NULL
);

-- Table products
CREATE TABLE IF NOT EXISTS products (
  id SERIAL PRIMARY KEY,
  name char(50) NOT NULL
);

-- Table orders
CREATE TABLE IF NOT EXISTS orders (
  id SERIAL PRIMARY KEY,
  customer_id INT NOT NULL,
  product_id INT NOT NULL,
  order_date DATE NOT NULL,
  total INT NOT NULL,
  FOREIGN KEY (customer_id)
    REFERENCES customers (id),
  FOREIGN KEY (product_id)
    REFERENCES products (id)
);

-- add 3 customers
INSERT INTO customers (fullname) VALUES
  ('Prasetya Priyadi'),
  ('Andhika Ramadhan'),
  ('Panji Ridwan') ;

-- Change user 2 to Andhika Ramdhan
UPDATE customers 
SET fullname = 'Andhika Ramdhan'
WHERE customers.id = 2;

-- add 5 products
INSERT INTO products (name) VALUES 
  ('Ayam Geprek'),
  ('Ayam Madu'),
  ('Ikan Nila Bakar'),
  ('Sop Ikan Mas'),
  ('Gulai Ikan Kakap');

-- add 5 Orders
INSERT INTO orders (customer_id, product_id, order_date, total) VALUES 
  (1, 1, '2022-9-21', 13000),
  (1, 3, '2022-9-21', 15000),
  (2, 2, '2022-9-23', 19500),
  (3, 4, '2022-9-30', 14000),
  (2, 5, '2022-10-1', 27000);