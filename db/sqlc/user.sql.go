// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
	first_name,
	last_name,
	email,
	phone,
	password,
	access_level
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING id, first_name, last_name, email, phone, password, access_level, created_at, updated_at
`

type CreateUserParams struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Password    string `json:"password"`
	AccessLevel int32  `json:"access_level"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Phone,
		arg.Password,
		arg.AccessLevel,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Phone,
		&i.Password,
		&i.AccessLevel,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, first_name, last_name, email, phone, password, access_level, created_at, updated_at FROM users
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUser, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Phone,
		&i.Password,
		&i.AccessLevel,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET
  password = COALESCE($1, password),
  updated_at = COALESCE($2, updated_at),
  first_name = COALESCE($3, first_name),
  last_name = COALESCE($4, last_name),
  phone = COALESCE($5, phone),
  access_level = COALESCE($6, access_level)
WHERE
  email = $7
RETURNING id, first_name, last_name, email, phone, password, access_level, created_at, updated_at
`

type UpdateUserParams struct {
	Password    pgtype.Text        `json:"password"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at"`
	FirstName   pgtype.Text        `json:"first_name"`
	LastName    pgtype.Text        `json:"last_name"`
	Phone       pgtype.Text        `json:"phone"`
	AccessLevel pgtype.Int4        `json:"access_level"`
	Email       string             `json:"email"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.Password,
		arg.UpdatedAt,
		arg.FirstName,
		arg.LastName,
		arg.Phone,
		arg.AccessLevel,
		arg.Email,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Phone,
		&i.Password,
		&i.AccessLevel,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
