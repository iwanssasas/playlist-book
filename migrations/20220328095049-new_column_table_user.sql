-- +migrate Up
ALTER TABLE users ADD google_id VARCHAR(100) UNIQUE;
ALTER TABLE users ADD is_edited BOOL NOT NULL DEFAULT false;
    
-- +migrate Down
ALTER TABLE users DROP COLUMN google_id;
ALTER TABLE users DROP COLUMN is_edited;