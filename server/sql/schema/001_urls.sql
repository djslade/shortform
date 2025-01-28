-- +goose Up
CREATE TABLE urls (
    id TEXT PRIMARY KEY,
    dest TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    expired_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE urls;