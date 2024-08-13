-- name: GetShopper :one
SELECT
    sqlc.embed(users),
    sqlc.embed(carts)
FROM
    users
    JOIN carts ON carts.user_id = users.id
WHERE
    users.id = ?;

