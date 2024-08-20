-- name: CreateCartItem :one
INSERT INTO cart_items (id, item_id, cart_id, quantity)
    VALUES (uuid (), ?, ?, ?)
RETURNING
    *;

-- name: DeleteCartItem :exec
DELETE FROM cart_items
WHERE id = ?;

-- name: UpdateCartItemQuantity :one
UPDATE
    cart_items
SET
    quantity = ?
WHERE
    id = ?
RETURNING
    *;

-- name: GetCartItems :many
SELECT
    *
FROM
    cart_items
WHERE
    cart_id = ?;

