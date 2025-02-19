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

-- name: UpdateUser :one
UPDATE users 
SET password_hash=$1,updated_at=NOW() 
WHERE id=$1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id=$1;