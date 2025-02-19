-- +goose Up
CREATE TABLE urls (
    id TEXT PRIMARY KEY,
    destination TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    expired_at TIMESTAMP
);

-- +goose Down
DROP TABLE urls;