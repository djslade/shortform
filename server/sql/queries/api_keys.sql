-- name: CreateAPIKey :one
INSERT INTO api_keys (key, created_at, expired_at) VALUES ($1, NOW(), NULL) RETURNING *;

-- name: GetAPIKey :one
SELECT * FROM api_keys WHERE key=$1;

-- name: ExpireAPIKey :exec
UPDATE api_keys SET expired_at=NOW() WHERE key=$1;

-- name: DeleteExpiredAPIKeys :exec
DELETE FROM api_keys WHERE expired_at <= NOW();