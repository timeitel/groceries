-- name: CreateUser :one
INSERT INTO users (id, name)
    VALUES (uuid (), ?)
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
INSERT INTO users (id, name, is_admin)
    VALUES (uuid (), ?, 1)
RETURNING
    *;

