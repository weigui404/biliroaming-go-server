-- +migrate Up
ALTER TABLE access_keys 
ADD COLUMN IF NOT EXISTS client_type VARCHAR(32);

-- +migrate Down
ALTER TABLE access_keys 
DROP COLUMN IF EXISTS client_type;
