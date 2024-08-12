-- name: CreateProduct :one
INSERT INTO products (name, description)
    VALUES (?, ?)
RETURNING
    *;

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

-- name: CreateCart :one
INSERT INTO carts (user_id)
    VALUES (?)
RETURNING
    *;

-- name: AddCartItem :one
INSERT INTO cart_items (product_id, cart_id, quantity)
    VALUES (?, ?, ?)
RETURNING
    *;

