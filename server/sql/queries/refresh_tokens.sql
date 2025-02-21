-- name: CreateRefreshToken :one
INSERT INTO refresh_tokens (
    token,
    user_id, 
    provider, 
    created_at, 
    expired_at
) VALUES (
    $1,
    $2,
    $3,
    NOW(),
    $4
)
RETURNING *;

-- name: RevokeRefreshToken :exec
UPDATE refresh_tokens SET expired_at=NOW() WHERE token=$1;

-- name: DeleteRefreshToken :exec
DELETE FROM refresh_tokens WHERE token=$1;

-- name: DeleteExpiredRefreshTokens :exec
DELETE FROM refresh_tokens WHERE expired_at <= NOW();