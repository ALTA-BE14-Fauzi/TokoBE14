use tokobe14;

show tables;

drop table users;
drop table transaksis;
drop table items;
drop table transaksi_items;
DROP TABLE customers;

CREATE TABLE users (
  id int auto_increment primary key,
  nama varchar(255) NOT NULL,
  password varchar(255) NOT NULL,
  role int NOT NULL
);

CREATE TABLE transaksis (
  id int auto_increment PRIMARY KEY ,
  user_id int NOT NULL,
  customer_id int NOT NULL,
  create_date date NOT NULL,
  FOREIGN KEY(user_id) REFERENCES users(id),
  FOREIGN KEY(customer_id) REFERENCES customers(id)
);

CREATE TABLE items (
  id int auto_increment primary key,
  nama varchar(255) NOT NULL,
  stock int NOT NULL,
  user_id int NOT NULL,
);

CREATE TABLE customers(
  id int auto_increment primary key,
  nama varchar(255) NOT NULL
);

CREATE TABLE transaksi_items(
  transaction_id int,
  item_id int,
  FOREIGN KEY(transaction_id) REFERENCES transaksis(id),
  FOREIGN KEY(item_id) REFERENCES items(id)
);

ALTER TABLE items
DROP CONSTRAINT items_ibfk_1;

ALTER TABLE transaksis 
DROP CONSTRAINT transaksis_ibfk_1;

INSERT INTO users (nama, password, role) VALUES ('admin', 'admin', 1);
INSERT INTO users (nama, password, role) VALUES ('messi', 'messi123', 2);


INSERT INTO items (nama, stock) VALUES ('indomie', 30);
INSERT INTO items (nama, stock) VALUES ('sedaap', 20);
INSERT INTO items (nama, stock) VALUES ('sarimie', 10);
INSERT INTO items (nama, stock) VALUES ('pop mie', 20);

INSERT INTO customers (nama) VALUES ('bryant');

INSERT INTO transaksis (user_id, customer_id, create_date) VALUES (2, 1, now());
INSERT INTO transaksis (user_id, customer_id, create_date) VALUES (2, 2, now());

INSERT INTO transaksi_items (transaction_id, item_id) VALUES (1, 1);
INSERT INTO transaksi_items (transaction_id, item_id) VALUES (1, 2);
INSERT INTO transaksi_items (transaction_id, item_id) VALUES (1, 3);

SELECT t.id,u.nama, c.nama,i.nama,COUNT(*), t.create_date 
FROM transaksi_items t2
LEFT JOIN transaksis t ON t.id = transaction_id 
JOIN customers c ON c.id = customer_id 
JOIN items i ON i.id = item_id 	
jOIN users u ON u.id = t.user_id
WHERE t.id = 1 OR t.id = 2 AND i.nama  = "betadine"
group by t.id,c.nama,i.nama,t.create_date
having COUNT(*) >= 1;

SELECT t.id, u.nama, c.nama, i.nama, COUNT(*), t.create_date  
FROM transaksi_items t2
LEFT JOIN transaksis t ON t.id = t2.transaction_id
LEFT JOIN customers c ON c.id = t.customer_id
LEFT JOIN items i ON i.id = t2.item_id
left JOIN users u ON u.id = t.user_id
WHERE t.id = 7
group by t.id, i.nama;



SELECT * FROM transaksi_items
JOIN transaksis t ON t.id = transaction_id
JOIN customers c ON c.id = customer_id
;

SELECT t.id, u.nama,c.nama,create_date FROM transaksis t
JOIN users u ON u.id = t.user_id
JOIN customers c ON c.id = t.customer_id;

SELECT t.transaction_id, i.nama, t2.create_date  FROM transaksi_items t
JOIN items i ON i.id = t.item_id
JOIN transaksis t2
WHERE i.nama = "bango" 
order by i.nama;

SELECT *
FROM items i
LEFT JOIN users u ON u.id = i.id;

SELECT i.id , i.nama,i.stock, u.nama FROM items i
LEFT JOIN users u ON u.id = i.user_id;

SELECT t.id, u.nama, c.nama, create_date
FROM transaksis t
LEFT JOIN users u ON u.id = user_id
JOIN customers c ON c.id = customer_id;
