-- name: CreateURL :one
INSERT INTO urls (id, created_at, updated_at, expired_at, dest)
VALUES (
    $1,
    NOW(),
    NOW(),
    $2,
    $3
)
RETURNING *;

-- name: CheckForURLWithID :one
SELECT COUNT(*) FROM urls WHERE id=$1;

-- name: GetURLByID :one
SELECT * FROM urls WHERE id=$1;