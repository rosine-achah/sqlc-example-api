-- name: CreateOrder :one
INSERT INTO orders (
    customer_name,
    customer_phone,
    total_amount,
    currency,
    status
)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetOrderByID :one
SELECT * FROM orders WHERE id = $1;

-- name: UpdateOrderStatus :one
UPDATE orders
SET status = $2
WHERE id = $1
RETURNING *;
