-- name: CreateOrder :one
INSERT INTO "orders" (customer_name) VALUES ($1)
RETURNING *;

-- name: GetOrder :one
SELECT * FROM orders WHERE id = $1;

-- name: GetOrders :many
SELECT * FROM orders;

-- name: UpdateOrder :one
UPDATE orders SET customer_name = $1 WHERE id = $2 RETURNING *;


-- name: DeleteOrder :exec
DELETE FROM orders WHERE id = $1;

-- name: CreateItem :one
INSERT INTO items (
    code,
    description,
    quantity,
    order_id
) VALUES ($1, $2, $3, $4) RETURNING *;


-- name: GetItemByOrderID :many
SELECT * FROM items WHERE order_id = $1;


-- name: DeleteItemByOrderID :exec
DELETE FROM items WHERE order_id = $1;