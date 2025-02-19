-- name: CreateClick :one
INSERT INTO clicks (
    id,
    url_id, 
    created_at, 
    continent, 
    country, 
    region, 
    city, 
    lat, 
    lon, 
    timezone, 
    currency, 
    referral_url, 
    device, 
    is_proxy, 
    isp
)
VALUES (
    GEN_RANDOM_UUID(),
    $1,
    NOW(),
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10,
    $11,
    $12,
    $13
)
RETURNING *;

-- name: GetClickByID :one
SELECT * FROM clicks WHERE id=$1;

-- name: GetClicksByURLID :many
SELECT * FROM clicks WHERE url_id=$1;