-- +migrate Up
ALTER TABLE
    users
MODIFY
    role_id int(3) NOT NULL;
    
-- +migrate Down
ALTER TABLE
    users
MODIFY
    role_id int(3);