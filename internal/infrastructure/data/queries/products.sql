-- name: CreateProduct :one
INSERT INTO products (name, description)
    VALUES (?, ?)
RETURNING
    *;

-- name: GetProducts :many
SELECT
    *
FROM
    products;

