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

-- name: GetShopper :one
SELECT
    u.name,
    u.is_admin AS isAdmin,
    c.name
FROM
    users u
    JOIN carts c ON c.id = u.active_cart_id
    -- join and get cart items
WHERE
    users.id = 1;

