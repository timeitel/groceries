-- name: CreateCart :one
INSERT INTO carts (user_id, name)
    VALUES (?, ?)
RETURNING
    *;

