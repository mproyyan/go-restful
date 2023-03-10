// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: user.sql

package repositories

import (
	"context"

	"github.com/lib/pq"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users(
    name,
    email
) VALUES (
    $1, $2
) RETURNING id, name, email, created_at
`

type CreateUserParams struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (q *Queries) CreateUser(ctx context.Context, db DBTX, arg CreateUserParams) (User, error) {
	row := db.QueryRowContext(ctx, createUser, arg.Name, arg.Email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.CreatedAt,
	)
	return i, err
}

const getBatchUsers = `-- name: GetBatchUsers :many
SELECT id, name, email, created_at FROM users
WHERE id = ANY($1::BIGINT[])
`

func (q *Queries) GetBatchUsers(ctx context.Context, db DBTX, ids []int64) ([]User, error) {
	rows, err := db.QueryContext(ctx, getBatchUsers, pq.Array(ids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUser = `-- name: GetUser :one
SELECT id, name, email, created_at FROM users 
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, db DBTX, id int64) (User, error) {
	row := db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.CreatedAt,
	)
	return i, err
}
