// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
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
RETURNING id, email, password_hash, created_at, updated_at
`

type CreateUserParams struct {
	Email        string
	PasswordHash sql.NullString
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Email, arg.PasswordHash)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE id=$1
`

func (q *Queries) DeleteUser(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, email, password_hash, created_at, updated_at FROM users WHERE email=$1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, email, password_hash, created_at, updated_at FROM users WHERE id=$1
`

func (q *Queries) GetUserByID(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserFromRefreshToken = `-- name: GetUserFromRefreshToken :one
SELECT users.id, users.email, users.password_hash, users.created_at, users.updated_at FROM users
JOIN refresh_tokens ON users.id=refresh_tokens.user_id
WHERE refresh_tokens.token=$1
AND expires_at > NOW()
`

func (q *Queries) GetUserFromRefreshToken(ctx context.Context, token string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserFromRefreshToken, token)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users 
SET password_hash=$1,updated_at=NOW() 
WHERE id=$1
RETURNING id, email, password_hash, created_at, updated_at
`

func (q *Queries) UpdateUser(ctx context.Context, passwordHash sql.NullString) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser, passwordHash)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
