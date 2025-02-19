-- +goose up
CREATE TABLE users (
    id UUID PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

ALTER TABLE urls
    ADD user_id UUID REFERENCES users(id) ON DELETE CASCADE 
;

-- +goose Down
ALTER TABLE urls
    DROP COLUMN user_id
;

DROP TABLE users;