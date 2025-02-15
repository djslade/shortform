-- +goose Up
CREATE TABLE clicks (
    id UUID PRIMARY KEY,
    url_id TEXT REFERENCES urls(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL,
    country_code
);

-- +goose Down
DROP TABLE clicks;