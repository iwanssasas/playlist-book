-- +migrate Up
CREATE TABLE test (
    id int(2) AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100)
);
-- +migrate Down
DROP TABLE test;