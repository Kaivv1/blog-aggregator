-- +goose Up
ALTER TABLE users
ADD COLUMN api_key VARCHAR(64) UNIQUE NOT NULL DEFAULT (
  encode(sha256(random()::text::bytea), 'hex')
); 

CREATE INDEX api_key_idx ON users(api_key);

-- +goose Down
DROP INDEX api_key_idx;

ALTER TABLE users
DROP COLUMN api_key;
