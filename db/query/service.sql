-- name: CreateService :one
INSERT INTO services (
    name,
    price
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetServiceById :one
SELECT * FROM services
WHERE id = $1
LIMIT 1;

-- name: GetServiceByName :one
SELECT * FROM services
WHERE name = $1
LIMIT 1;

-- name: ListServices :many
SELECT * FROM services
ORDER BY name;

-- name: DeleteService :exec
DELETE FROM services WHERE id = $1;
