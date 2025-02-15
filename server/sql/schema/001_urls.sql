-- +goose Up
CREATE TABLE urls (
    id TEXT PRIMARY KEY,
    dest TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    disabled_at TIMESTAMP
);

-- +goose Down
DROP TABLE urls;