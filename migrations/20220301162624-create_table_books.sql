-- +migrate Up
CREATE TABLE books (
    ID INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    quantity INT(10)
);
-- +migrate Down
DROP TABLE IF EXISTS books;