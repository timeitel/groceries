-- name: CreateCartItem :one
INSERT INTO cart_items (item_id, cart_id, quantity)
    VALUES (?, ?, ?)
RETURNING
    *;

-- name: DeleteCartItem :exec
INSERT INTO cart_items (cart_id, item_id)
    VALUES (?, ?)
RETURNING
    *;

-- name: UpdateCartItemQuantity :one
UPDATE
    cart_items
SET
    "quantity" = ?
WHERE
    item_id = ?
RETURNING
    *;

-- name: GetCartItems :many
SELECT
    *
FROM
    cart_items
WHERE
    cart_id = ?;

