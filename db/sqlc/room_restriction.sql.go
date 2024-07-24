// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: room_restriction.sql

package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

const createRoomRestriction = `-- name: CreateRoomRestriction :one
INSERT INTO room_restrictions (
	start_date,
	end_date,
	room_id,
	reservation_id,
	restriction_id
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING id, start_date, end_date, room_id, reservation_id, restriction_id, created_at, updated_at
`

type CreateRoomRestrictionParams struct {
	StartDate     time.Time   `json:"start_date"`
	EndDate       time.Time   `json:"end_date"`
	RoomID        int32       `json:"room_id"`
	ReservationID pgtype.Int4 `json:"reservation_id"`
	RestrictionID int32       `json:"restriction_id"`
}

func (q *Queries) CreateRoomRestriction(ctx context.Context, arg CreateRoomRestrictionParams) (RoomRestriction, error) {
	row := q.db.QueryRow(ctx, createRoomRestriction,
		arg.StartDate,
		arg.EndDate,
		arg.RoomID,
		arg.ReservationID,
		arg.RestrictionID,
	)
	var i RoomRestriction
	err := row.Scan(
		&i.ID,
		&i.StartDate,
		&i.EndDate,
		&i.RoomID,
		&i.ReservationID,
		&i.RestrictionID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteRoomRestriction = `-- name: DeleteRoomRestriction :exec
DELETE FROM room_restrictions
WHERE id = $1
`

func (q *Queries) DeleteRoomRestriction(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteRoomRestriction, id)
	return err
}

const getRestrictionsForRoomByDate = `-- name: GetRestrictionsForRoomByDate :many
select id, start_date, end_date, room_id, reservation_id, restriction_id, created_at, updated_at
from room_restrictions where $1 < end_date and $2 >= start_date
and room_id = $3
ORDER BY id
LIMIT $4
OFFSET $5
`

type GetRestrictionsForRoomByDateParams struct {
	EndDate   time.Time `json:"end_date"`
	StartDate time.Time `json:"start_date"`
	RoomID    int32     `json:"room_id"`
	Limit     int32     `json:"limit"`
	Offset    int32     `json:"offset"`
}

func (q *Queries) GetRestrictionsForRoomByDate(ctx context.Context, arg GetRestrictionsForRoomByDateParams) ([]RoomRestriction, error) {
	rows, err := q.db.Query(ctx, getRestrictionsForRoomByDate,
		arg.EndDate,
		arg.StartDate,
		arg.RoomID,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []RoomRestriction{}
	for rows.Next() {
		var i RoomRestriction
		if err := rows.Scan(
			&i.ID,
			&i.StartDate,
			&i.EndDate,
			&i.RoomID,
			&i.ReservationID,
			&i.RestrictionID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRoomRestriction = `-- name: GetRoomRestriction :one
SELECT id, start_date, end_date, room_id, reservation_id, restriction_id, created_at, updated_at FROM room_restrictions
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetRoomRestriction(ctx context.Context, id int32) (RoomRestriction, error) {
	row := q.db.QueryRow(ctx, getRoomRestriction, id)
	var i RoomRestriction
	err := row.Scan(
		&i.ID,
		&i.StartDate,
		&i.EndDate,
		&i.RoomID,
		&i.ReservationID,
		&i.RestrictionID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getRoomRestrictionForUpdate = `-- name: GetRoomRestrictionForUpdate :one
SELECT id, start_date, end_date, room_id, reservation_id, restriction_id, created_at, updated_at FROM room_restrictions
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetRoomRestrictionForUpdate(ctx context.Context, id int32) (RoomRestriction, error) {
	row := q.db.QueryRow(ctx, getRoomRestrictionForUpdate, id)
	var i RoomRestriction
	err := row.Scan(
		&i.ID,
		&i.StartDate,
		&i.EndDate,
		&i.RoomID,
		&i.ReservationID,
		&i.RestrictionID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listRoomRestrictions = `-- name: ListRoomRestrictions :many
SELECT id, start_date, end_date, room_id, reservation_id, restriction_id, created_at, updated_at FROM room_restrictions
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListRoomRestrictionsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListRoomRestrictions(ctx context.Context, arg ListRoomRestrictionsParams) ([]RoomRestriction, error) {
	rows, err := q.db.Query(ctx, listRoomRestrictions, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []RoomRestriction{}
	for rows.Next() {
		var i RoomRestriction
		if err := rows.Scan(
			&i.ID,
			&i.StartDate,
			&i.EndDate,
			&i.RoomID,
			&i.ReservationID,
			&i.RestrictionID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const searchAvailabilityByDatesByRoomID = `-- name: SearchAvailabilityByDatesByRoomID :one
select count(id)
from room_restrictions
where room_id = $1
and $2 < end_date and $3 >= start_date
`

type SearchAvailabilityByDatesByRoomIDParams struct {
	RoomID    int32     `json:"room_id"`
	EndDate   time.Time `json:"end_date"`
	StartDate time.Time `json:"start_date"`
}

func (q *Queries) SearchAvailabilityByDatesByRoomID(ctx context.Context, arg SearchAvailabilityByDatesByRoomIDParams) (int64, error) {
	row := q.db.QueryRow(ctx, searchAvailabilityByDatesByRoomID, arg.RoomID, arg.EndDate, arg.StartDate)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const updateRoomRestriction = `-- name: UpdateRoomRestriction :one
UPDATE room_restrictions
SET
  updated_at = COALESCE($1, updated_at),
  start_date = COALESCE($2, start_date),
  end_date = COALESCE($3, end_date),
  room_id = COALESCE($4, room_id),
  reservation_id = COALESCE($5, reservation_id),
  restriction_id = COALESCE($6, restriction_id)
WHERE
id = $7
RETURNING id, start_date, end_date, room_id, reservation_id, restriction_id, created_at, updated_at
`

type UpdateRoomRestrictionParams struct {
	UpdatedAt     pgtype.Timestamptz `json:"updated_at"`
	StartDate     pgtype.Date        `json:"start_date"`
	EndDate       pgtype.Date        `json:"end_date"`
	RoomID        pgtype.Int4        `json:"room_id"`
	ReservationID pgtype.Int4        `json:"reservation_id"`
	RestrictionID pgtype.Int4        `json:"restriction_id"`
	ID            int32              `json:"id"`
}

func (q *Queries) UpdateRoomRestriction(ctx context.Context, arg UpdateRoomRestrictionParams) (RoomRestriction, error) {
	row := q.db.QueryRow(ctx, updateRoomRestriction,
		arg.UpdatedAt,
		arg.StartDate,
		arg.EndDate,
		arg.RoomID,
		arg.ReservationID,
		arg.RestrictionID,
		arg.ID,
	)
	var i RoomRestriction
	err := row.Scan(
		&i.ID,
		&i.StartDate,
		&i.EndDate,
		&i.RoomID,
		&i.ReservationID,
		&i.RestrictionID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
