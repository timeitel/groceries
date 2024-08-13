-- name: CreateUser :one
INSERT INTO users (name)
    VALUES (?)
RETURNING
    *;

-- name: GetUser :one
SELECT
    *
FROM
    users
WHERE
    id = ?;

-- name: CreateAdminUser :one
INSERT INTO users (name, is_admin)
    VALUES (?, 1)
RETURNING
    *;

