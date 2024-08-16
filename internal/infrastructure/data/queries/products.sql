-- name: CreateProduct :one
INSERT INTO products (name, description)
    VALUES (?, ?)
RETURNING
    *;

-- name: GetProduct :one
SELECT
    *
FROM
    products
WHERE
    id = ?;

-- name: GetProducts :many
SELECT
    *
FROM
    products;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = ?;

-- name: UpdateProduct :exec
UPDATE
    products
SET
    name = ?,
    description = ?
WHERE
    id = ?;

