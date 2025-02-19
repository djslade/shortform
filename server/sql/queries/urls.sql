-- name: CreateURL :one
INSERT INTO urls (id, user_id, created_at, updated_at, expired_at, destination)
VALUES (
    $1,
    $2,
    NOW(),
    NOW(),
    $3,
    $4
)
RETURNING *;

-- name: CheckForURLWithID :one
SELECT COUNT(*) FROM urls WHERE id=$1;

-- name: GetURLByID :one
SELECT * FROM urls WHERE id=$1;

-- name: GetURLsByUserID :many
SELECT * FROM urls WHERE user_id=$1;

-- name: GetURLsByDestination :many
SELECT * FROM urls WHERE user_id=$1 AND destination=$2;

-- name: UpdateURL :one
UPDATE urls 
SET expired_at=$1,destination=$2,updated_at=NOW() 
WHERE id=$3
RETURNING *;

-- name: DeleteURL :exec
DELETE FROM urls WHERE id=$1;