-- +goose Up
CREATE TABLE urls (
    id TEXT PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    expired_at TIMESTAMP
);

-- +goose Down
DROP TABLE urls;