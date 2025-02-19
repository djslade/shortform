-- +goose Up
CREATE TABLE clicks (
    id UUID PRIMARY KEY,
    url_id TEXT REFERENCES urls(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL,
    continent TEXT,
    country TEXT,
    region TEXT,
    city TEXT,
    lat FLOAT,
    lon FLOAT,
    timezone TEXT,
    currency TEXT,
    referral_url TEXT,
    device TEXT,
    is_proxy BOOLEAN,
    isp TEXT

);

-- +goose Down
DROP TABLE clicks;