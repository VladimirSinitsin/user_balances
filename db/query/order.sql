-- name: CreateOrder :one
INSERT INTO orders (
    id_account,
    id_service,
    price_service,
    status
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetOrderById :one
SELECT * FROM orders
WHERE id = $1
LIMIT 1;

-- name: ListOrdersByServiceId :many
SELECT * FROM orders
WHERE id_service = $1;

-- name: UpdateOrder :one
UPDATE orders
SET status = $2
WHERE id = $1
RETURNING *;

-- name: DeleteOrder :exec
DELETE FROM orders WHERE id = $1;
