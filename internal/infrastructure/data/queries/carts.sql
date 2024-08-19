-- name: CreateCart :one
INSERT INTO carts (user_id, name)
    VALUES (?, ?)
RETURNING
    *;

-- name: GetCart :one
SELECT
    *
FROM
    carts
WHERE
    id = ?;

