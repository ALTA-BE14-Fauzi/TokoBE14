CREATE DATABASE tokobe14;

USE tokobe14;

drop table users;
drop table transaksis;

CREATE TABLE users (
  id int auto_increment primary key,
  nama varchar(255) NOT NULL,
  password varchar(255) NOT NULL,
  role int NOT NULL
);

CREATE TABLE items (
  id int auto_increment primary key,
  nama varchar(255) NOT NULL,
  stock int NOT NULL
);

CREATE TABLE customers(
  id int auto_increment primary key,
  nama varchar(255) NOT NULL
);

CREATE TABLE transaksis (
  id int auto_increment primary key,
  user_id int NOT NULL,
  item_id int NOT NULL,
  customer_id int NOT NULL,
  create_date date NOT NULL,
  FOREIGN KEY(user_id) REFERENCES users(id),
  FOREIGN KEY(item_id) REFERENCES items(id),
  FOREIGN KEY(customer_id) REFERENCES customers(id)
);

CREATE TABLE transaksi_items(
  transaction_id int NOT NULL,
  item_id int NOT NULL,
  qty int NOT NULL,
  PRIMARY KEY(transaction_id, item_id),
  FOREIGN KEY(transaction_id) REFERENCES transaksis(id),
  FOREIGN KEY(item_id) REFERENCES items(id)
);

INSERT INTO users (nama, password, role) VALUES ('admin', 'admin', 1);
INSERT INTO users (nama, password, role) VALUES ('messi', 'messi123', 2);
INSERT INTO users (nama, password, role) VALUES ('voldemort', '123', 2);

INSERT INTO items (nama, stock) VALUES ('indomie', 30);

INSERT INTO customers (nama) VALUES ('bryant');

INSERT INTO transaksis (user_id, item_id, customer_id, create_date) VALUES (2, 1, 1, '2022-12-19');

INSERT INTO transaksi_items (transaction_id, item_id, qty) VALUES (1, 1, 1);

SELECT * FROM users;
SELECT * FROM items;
SELECT * FROM customers;
SELECT * FROM transaksis;