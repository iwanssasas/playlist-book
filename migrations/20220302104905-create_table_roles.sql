
-- +migrate Up
CREATE TABLE roles (
    id INT(3) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    description TEXT
);
-- +migrate Down
DROP TABLE IF EXISTS roles;
