-- name: CheckItem :one
SELECT exists(
  select 1 from Inventory where app_id = $1 AND asset_id = $2 LIMIT 1
);

-- name: InsertItem :one
INSERT INTO Inventory (
  bot_id,
  app_id,
  asset_id,
  class_id,
  instance_id
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: RemoveItem :exec
DELETE FROM Inventory
WHERE app_id = $1 AND asset_id = $2;