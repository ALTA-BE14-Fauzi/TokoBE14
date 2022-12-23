CREATE DATABASE tokobe14;

USE tokobe14;

drop table users;
drop table transaksis;
drop table customers;
drop table transaksi_items;
drop table items;

CREATE TABLE users (
  id int auto_increment primary key,
  nama varchar(255) NOT NULL,
  password varchar(255) NOT NULL,
  role int NOT NULL
);

CREATE TABLE items (
  id int auto_increment primary key,
  nama varchar(255) NOT NULL,
  stock int NOT NULL,
  user_id int NOT NULL,
  FOREIGN KEY(user_id) REFERENCES users(id)
);

CREATE TABLE customers(
  id int auto_increment primary key,
  nama varchar(255) NOT NULL
);

CREATE TABLE transaksis (
  id int auto_increment PRIMARY KEY ,
  user_id int NOT NULL,
  customer_id int NOT NULL,
  create_date date NOT NULL,
  FOREIGN KEY(user_id) REFERENCES users(id),
  FOREIGN KEY(customer_id) REFERENCES customers(id)
);

CREATE TABLE transaksi_items(
  transaction_id int NOT NULL,
  item_id int NOT NULL,
  FOREIGN KEY(transaction_id) REFERENCES transaksis(id),
  FOREIGN KEY(item_id) REFERENCES items(id)
);

INSERT INTO users (nama, password, role) VALUES ('admin', 'admin', 1);
INSERT INTO users (nama, password, role) VALUES ('messi', 'messi123', 2);
INSERT INTO users (nama, password, role) VALUES ('voldemort', '123', 2);

INSERT INTO items (nama, stock) VALUES ('indomie', 30);

INSERT INTO customers (nama) VALUES ('bryant');

INSERT INTO transaksis (user_id, customer_id, create_date) VALUES (2, 1, now());
INSERT INTO transaksis (user_id, customer_id, create_date) VALUES (2, 2, now());

INSERT INTO transaksi_items (transaction_id, item_id) VALUES (1, 1);
INSERT INTO transaksi_items (transaction_id, item_id) VALUES (1, 2);
INSERT INTO transaksi_items (transaction_id, item_id) VALUES (1, 3);

SELECT * FROM users;
SELECT * FROM items;
SELECT * FROM customers;
SELECT * FROM transaksis;
SELECT * FROM transaksi_items;

SELECT i.id , i.nama,i.stock,u.nama  
FROM items i
JOIN users u ON u.id = i.user_id;

SELECT t.id,u.nama "Kasir", c.nama "Customer",i.nama,COUNT(*) "Quantity",  t.create_date
FROM transaksi_items 
JOIN transaksis t ON t.id = transaction_id
JOIN customers c ON c.id = customer_id
JOIN items i ON i.id = item_id
jOIN users u ON u.id = t.user_id
WHERE t.id = 2
group by t.id,c.nama,i.nama,t.create_date 
having COUNT(*) >= 1;

SELECT t.id,u.nama "Nama_Kasir",c.nama "Nama_Customer",create_date "Tanggal_Pembelian"
FROM transaksis t
JOIN users u ON u.id = user_id
JOIN customers c ON c.id = customer_id
ORDER BY t.id;

SELECT t.id FROM transaksis t order by t.id asc;
