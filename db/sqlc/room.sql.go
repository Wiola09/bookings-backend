// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: room.sql

package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

const createRoom = `-- name: CreateRoom :one
INSERT INTO rooms (
	room_name_sr,
	room_name_en,
	room_name_bg,
	room_short_des_sr,
	room_short_des_en,
	room_short_des_bg,
	room_description_sr,
	room_description_en,
	room_description_bg,
	room_pictures_folder,
	room_guest_number,
	room_price_en
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
) RETURNING id, room_name_sr, room_name_en, room_name_bg, room_short_des_sr, room_short_des_en, room_short_des_bg, room_description_sr, room_description_en, room_description_bg, room_pictures_folder, room_guest_number, room_price_en, created_at, updated_at
`

type CreateRoomParams struct {
	RoomNameSr         string `json:"room_name_sr"`
	RoomNameEn         string `json:"room_name_en"`
	RoomNameBg         string `json:"room_name_bg"`
	RoomShortDesSr     string `json:"room_short_des_sr"`
	RoomShortDesEn     string `json:"room_short_des_en"`
	RoomShortDesBg     string `json:"room_short_des_bg"`
	RoomDescriptionSr  string `json:"room_description_sr"`
	RoomDescriptionEn  string `json:"room_description_en"`
	RoomDescriptionBg  string `json:"room_description_bg"`
	RoomPicturesFolder string `json:"room_pictures_folder"`
	RoomGuestNumber    int32  `json:"room_guest_number"`
	RoomPriceEn        int32  `json:"room_price_en"`
}

func (q *Queries) CreateRoom(ctx context.Context, arg CreateRoomParams) (Room, error) {
	row := q.db.QueryRow(ctx, createRoom,
		arg.RoomNameSr,
		arg.RoomNameEn,
		arg.RoomNameBg,
		arg.RoomShortDesSr,
		arg.RoomShortDesEn,
		arg.RoomShortDesBg,
		arg.RoomDescriptionSr,
		arg.RoomDescriptionEn,
		arg.RoomDescriptionBg,
		arg.RoomPicturesFolder,
		arg.RoomGuestNumber,
		arg.RoomPriceEn,
	)
	var i Room
	err := row.Scan(
		&i.ID,
		&i.RoomNameSr,
		&i.RoomNameEn,
		&i.RoomNameBg,
		&i.RoomShortDesSr,
		&i.RoomShortDesEn,
		&i.RoomShortDesBg,
		&i.RoomDescriptionSr,
		&i.RoomDescriptionEn,
		&i.RoomDescriptionBg,
		&i.RoomPicturesFolder,
		&i.RoomGuestNumber,
		&i.RoomPriceEn,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteRoom = `-- name: DeleteRoom :exec
DELETE FROM rooms
WHERE id = $1
`

func (q *Queries) DeleteRoom(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteRoom, id)
	return err
}

const getRoom = `-- name: GetRoom :one
SELECT id, room_name_sr, room_name_en, room_name_bg, room_short_des_sr, room_short_des_en, room_short_des_bg, room_description_sr, room_description_en, room_description_bg, room_pictures_folder, room_guest_number, room_price_en, created_at, updated_at FROM rooms
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetRoom(ctx context.Context, id int32) (Room, error) {
	row := q.db.QueryRow(ctx, getRoom, id)
	var i Room
	err := row.Scan(
		&i.ID,
		&i.RoomNameSr,
		&i.RoomNameEn,
		&i.RoomNameBg,
		&i.RoomShortDesSr,
		&i.RoomShortDesEn,
		&i.RoomShortDesBg,
		&i.RoomDescriptionSr,
		&i.RoomDescriptionEn,
		&i.RoomDescriptionBg,
		&i.RoomPicturesFolder,
		&i.RoomGuestNumber,
		&i.RoomPriceEn,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getRoomForUpdate = `-- name: GetRoomForUpdate :one
SELECT id, room_name_sr, room_name_en, room_name_bg, room_short_des_sr, room_short_des_en, room_short_des_bg, room_description_sr, room_description_en, room_description_bg, room_pictures_folder, room_guest_number, room_price_en, created_at, updated_at FROM rooms
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetRoomForUpdate(ctx context.Context, id int32) (Room, error) {
	row := q.db.QueryRow(ctx, getRoomForUpdate, id)
	var i Room
	err := row.Scan(
		&i.ID,
		&i.RoomNameSr,
		&i.RoomNameEn,
		&i.RoomNameBg,
		&i.RoomShortDesSr,
		&i.RoomShortDesEn,
		&i.RoomShortDesBg,
		&i.RoomDescriptionSr,
		&i.RoomDescriptionEn,
		&i.RoomDescriptionBg,
		&i.RoomPicturesFolder,
		&i.RoomGuestNumber,
		&i.RoomPriceEn,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listRooms = `-- name: ListRooms :many
SELECT id, room_name_sr, room_name_en, room_name_bg, room_short_des_sr, room_short_des_en, room_short_des_bg, room_description_sr, room_description_en, room_description_bg, room_pictures_folder, room_guest_number, room_price_en, created_at, updated_at FROM rooms
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListRoomsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListRooms(ctx context.Context, arg ListRoomsParams) ([]Room, error) {
	rows, err := q.db.Query(ctx, listRooms, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Room{}
	for rows.Next() {
		var i Room
		if err := rows.Scan(
			&i.ID,
			&i.RoomNameSr,
			&i.RoomNameEn,
			&i.RoomNameBg,
			&i.RoomShortDesSr,
			&i.RoomShortDesEn,
			&i.RoomShortDesBg,
			&i.RoomDescriptionSr,
			&i.RoomDescriptionEn,
			&i.RoomDescriptionBg,
			&i.RoomPicturesFolder,
			&i.RoomGuestNumber,
			&i.RoomPriceEn,
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

const searchAvailabilityForAllRooms = `-- name: SearchAvailabilityForAllRooms :many
select r.id, r.room_name_sr, r.room_name_en, r.room_name_bg, r.room_short_des_sr, r.room_short_des_en, r.room_short_des_bg, r.room_description_sr, r.room_description_en, r.room_description_bg, r.room_pictures_folder, r.room_guest_number, r.room_price_en, r.created_at, r.updated_at from rooms r
where r.id not in 
		(select room_id from room_restrictions rr where $1 < rr.end_date and $2 >= rr.start_date)
ORDER BY r.id
LIMIT $3
OFFSET $4
`

type SearchAvailabilityForAllRoomsParams struct {
	EndDate   time.Time `json:"end_date"`
	StartDate time.Time `json:"start_date"`
	Limit     int32     `json:"limit"`
	Offset    int32     `json:"offset"`
}

func (q *Queries) SearchAvailabilityForAllRooms(ctx context.Context, arg SearchAvailabilityForAllRoomsParams) ([]Room, error) {
	rows, err := q.db.Query(ctx, searchAvailabilityForAllRooms,
		arg.EndDate,
		arg.StartDate,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Room{}
	for rows.Next() {
		var i Room
		if err := rows.Scan(
			&i.ID,
			&i.RoomNameSr,
			&i.RoomNameEn,
			&i.RoomNameBg,
			&i.RoomShortDesSr,
			&i.RoomShortDesEn,
			&i.RoomShortDesBg,
			&i.RoomDescriptionSr,
			&i.RoomDescriptionEn,
			&i.RoomDescriptionBg,
			&i.RoomPicturesFolder,
			&i.RoomGuestNumber,
			&i.RoomPriceEn,
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

const updateRoom = `-- name: UpdateRoom :one
UPDATE rooms
SET
  updated_at = COALESCE($1, updated_at),
  room_name_sr = COALESCE($2, room_name_sr),
  room_name_en = COALESCE($3, room_name_en),
  room_name_bg = COALESCE($4, room_name_bg),
  room_short_des_sr = COALESCE($5, room_short_des_sr),
  room_short_des_en = COALESCE($6, room_short_des_en),
  room_short_des_bg = COALESCE($7, room_short_des_bg),
  room_description_sr = COALESCE($8, room_description_sr),
  room_description_en = COALESCE($9, room_description_en),
  room_description_bg = COALESCE($10, room_description_bg),
  room_pictures_folder = COALESCE($11, room_pictures_folder),
  room_guest_number = COALESCE($12, room_guest_number),
  room_price_en = COALESCE($13, room_price_en)
WHERE
id = $14
RETURNING id, room_name_sr, room_name_en, room_name_bg, room_short_des_sr, room_short_des_en, room_short_des_bg, room_description_sr, room_description_en, room_description_bg, room_pictures_folder, room_guest_number, room_price_en, created_at, updated_at
`

type UpdateRoomParams struct {
	UpdatedAt          pgtype.Timestamptz `json:"updated_at"`
	RoomNameSr         pgtype.Text        `json:"room_name_sr"`
	RoomNameEn         pgtype.Text        `json:"room_name_en"`
	RoomNameBg         pgtype.Text        `json:"room_name_bg"`
	RoomShortDesSr     pgtype.Text        `json:"room_short_des_sr"`
	RoomShortDesEn     pgtype.Text        `json:"room_short_des_en"`
	RoomShortDesBg     pgtype.Text        `json:"room_short_des_bg"`
	RoomDescriptionSr  pgtype.Text        `json:"room_description_sr"`
	RoomDescriptionEn  pgtype.Text        `json:"room_description_en"`
	RoomDescriptionBg  pgtype.Text        `json:"room_description_bg"`
	RoomPicturesFolder pgtype.Text        `json:"room_pictures_folder"`
	RoomGuestNumber    pgtype.Int4        `json:"room_guest_number"`
	RoomPriceEn        pgtype.Int4        `json:"room_price_en"`
	ID                 int32              `json:"id"`
}

func (q *Queries) UpdateRoom(ctx context.Context, arg UpdateRoomParams) (Room, error) {
	row := q.db.QueryRow(ctx, updateRoom,
		arg.UpdatedAt,
		arg.RoomNameSr,
		arg.RoomNameEn,
		arg.RoomNameBg,
		arg.RoomShortDesSr,
		arg.RoomShortDesEn,
		arg.RoomShortDesBg,
		arg.RoomDescriptionSr,
		arg.RoomDescriptionEn,
		arg.RoomDescriptionBg,
		arg.RoomPicturesFolder,
		arg.RoomGuestNumber,
		arg.RoomPriceEn,
		arg.ID,
	)
	var i Room
	err := row.Scan(
		&i.ID,
		&i.RoomNameSr,
		&i.RoomNameEn,
		&i.RoomNameBg,
		&i.RoomShortDesSr,
		&i.RoomShortDesEn,
		&i.RoomShortDesBg,
		&i.RoomDescriptionSr,
		&i.RoomDescriptionEn,
		&i.RoomDescriptionBg,
		&i.RoomPicturesFolder,
		&i.RoomGuestNumber,
		&i.RoomPriceEn,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
