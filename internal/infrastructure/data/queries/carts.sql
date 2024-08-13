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

