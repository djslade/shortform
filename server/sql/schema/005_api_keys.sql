-- +goose Up
CREATE TABLE api_keys (
    key TEXT PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    expired_at TIMESTAMP
);

ALTER TABLE urls
ADD key_id TEXT REFERENCES api_keys(key) ON DELETE CASCADE;

-- +goose Down
ALTER TABLE urls
DROP COLUMN key_id;

DROP TABLE api_keys;