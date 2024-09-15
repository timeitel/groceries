-- name: CreateItem :one
INSERT INTO items (name, description)
    VALUES (?, ?)
RETURNING
    *;

-- name: GetItem :one
SELECT
    *
FROM
    items
WHERE
    id = ?;

-- name: GetItems :many
SELECT
    *
FROM
    items;

-- name: DeleteItem :exec
DELETE FROM items
WHERE id = ?;

-- name: UpdateItem :one
UPDATE
    items
SET
    name = ?,
    description = ?
WHERE
    id = ?
RETURNING
    *;

