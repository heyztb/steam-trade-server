-- name: InsertItem :exec
INSERT INTO Items (
  bot_id,
  app_id,
  asset_id,
  class_id,
  instance_id
) VALUES (
  ?, ?, ?, ?, ? 
);

-- name: RemoveItem :exec
DELETE FROM Items
WHERE app_id = ? AND asset_id = ?;

-- name: GetAllBots :many
SELECT * FROM Bots;

-- name: GetRandomBot :one
SELECT * FROM Bots ORDER BY random() LIMIT 1;

-- name: InsertBot :exec
INSERT INTO Bots (
  username,
  passwd,
  shared_secret,
  identity_secret,
  api_key
) VALUES (
  ?, ?, ?, ?, ?
);