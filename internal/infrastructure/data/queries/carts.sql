-- name: CreateCart :one
INSERT INTO carts (id, user_id, name)
    VALUES (uuid (), ?, ?)
RETURNING
    *;

-- name: GetCarts :one
SELECT
    *
FROM
    carts
WHERE
    user_id = ?;

-- name: GetCart :one
SELECT
    *
FROM
    carts
WHERE
    user_id = ?
LIMIT 1;

