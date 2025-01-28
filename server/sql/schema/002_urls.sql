-- +goose Up
ALTER TABLE urls
ADD dest TEXT NOT NULL;

-- +goose Down
ALTER TABLE urls
DROP COLUMN dest;