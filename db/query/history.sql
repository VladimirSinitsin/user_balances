-- name: CreateHistory :one
INSERT INTO history (
    id_account,
    amount,
    comment
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetHistoryById :one
SELECT * FROM history
WHERE id = $1
LIMIT 1;

-- name: ListHistoryByAccountIdDataSort :many
SELECT * FROM history
WHERE id_account = $1
ORDER BY created_at
LIMIT $2
OFFSET $3;

-- name: ListHistoryByAccountIdAmountSort :many
SELECT * FROM history
WHERE id_account = $1
ORDER BY amount
LIMIT $2
OFFSET $3;

-- name: DeleteHistory :exec
DELETE FROM history WHERE id = $1;
