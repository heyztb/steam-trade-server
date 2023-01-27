-- name: CheckItem :one
SELECT exists(
  select 1 from Items where app_id = $1 AND asset_id = $2 LIMIT 1
);

-- name: InsertItem :exec
INSERT INTO Items (
  bot_id,
  app_id,
  asset_id,
  class_id,
  instance_id
) VALUES (
  $1, $2, $3, $4, $5
);

-- name: RemoveItem :exec
DELETE FROM Items
WHERE app_id = $1 AND asset_id = $2;

-- name: GetAllBots :many
SELECT * FROM Bots;

-- name: InsertBot :exec
INSERT INTO Bots (
  username,
  passwd,
  shared_secret,
  identity_secret
) VALUES (
  $1, $2, $3, $4
);
