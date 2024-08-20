-- name: CreateCart :one
INSERT INTO carts (id, user_id, name)
    VALUES (uuid (), ?, ?)
RETURNING
    *;

-- name: GetCart :one
SELECT
    *
FROM
    carts
WHERE
    id = ?;

