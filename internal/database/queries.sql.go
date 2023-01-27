// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: queries.sql

package database

import (
	"context"
)

const checkItem = `-- name: CheckItem :one
SELECT exists(
  select 1 from Items where app_id = $1 AND asset_id = $2 LIMIT 1
)
`

type CheckItemParams struct {
	AppID   int64
	AssetID int64
}

func (q *Queries) CheckItem(ctx context.Context, arg CheckItemParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkItem, arg.AppID, arg.AssetID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const getAllBots = `-- name: GetAllBots :many
SELECT id, username, passwd, shared_secret, identity_secret FROM Bots
`

func (q *Queries) GetAllBots(ctx context.Context) ([]Bot, error) {
	rows, err := q.db.QueryContext(ctx, getAllBots)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Bot
	for rows.Next() {
		var i Bot
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Passwd,
			&i.SharedSecret,
			&i.IdentitySecret,
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

const insertBot = `-- name: InsertBot :exec
INSERT INTO Bots (
  username,
  passwd,
  shared_secret,
  identity_secret
) VALUES (
  $1, $2, $3, $4
)
`

type InsertBotParams struct {
	Username       string
	Passwd         string
	SharedSecret   string
	IdentitySecret string
}

func (q *Queries) InsertBot(ctx context.Context, arg InsertBotParams) error {
	_, err := q.db.ExecContext(ctx, insertBot,
		arg.Username,
		arg.Passwd,
		arg.SharedSecret,
		arg.IdentitySecret,
	)
	return err
}

const insertItem = `-- name: InsertItem :exec
INSERT INTO Items (
  bot_id,
  app_id,
  asset_id,
  class_id,
  instance_id
) VALUES (
  $1, $2, $3, $4, $5
)
`

type InsertItemParams struct {
	BotID      int64
	AppID      int64
	AssetID    int64
	ClassID    int64
	InstanceID int64
}

func (q *Queries) InsertItem(ctx context.Context, arg InsertItemParams) error {
	_, err := q.db.ExecContext(ctx, insertItem,
		arg.BotID,
		arg.AppID,
		arg.AssetID,
		arg.ClassID,
		arg.InstanceID,
	)
	return err
}

const removeItem = `-- name: RemoveItem :exec
DELETE FROM Items
WHERE app_id = $1 AND asset_id = $2
`

type RemoveItemParams struct {
	AppID   int64
	AssetID int64
}

func (q *Queries) RemoveItem(ctx context.Context, arg RemoveItemParams) error {
	_, err := q.db.ExecContext(ctx, removeItem, arg.AppID, arg.AssetID)
	return err
}
