-- name: CreateUser :one
INSERT INTO users (
    id,
    email,
    password_hash,
    created_at,
    updated_at
)
VALUES (
    GEN_RANDOM_UUID(),
    $1,
    $2,
    NOW(),
    NOW()
)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users WHERE id=$1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email=$1;

-- name: GetUserFromRefreshToken :one
SELECT users.* FROM users
JOIN refresh_tokens ON users.id=refresh_tokens.user_id
WHERE refresh_tokens.token=$1
AND expires_at > NOW();

-- name: UpdateUser :exec
UPDATE users 
SET password_hash=$1,updated_at=NOW() 
WHERE id=$1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id=$1;