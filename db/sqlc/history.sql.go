// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: history.sql

package db

import (
	"context"
)

const createHistory = `-- name: CreateHistory :one
INSERT INTO history (
    id_account,
    amount,
    comment
) VALUES (
    $1, $2, $3
) RETURNING id, id_account, amount, comment, created_at
`

type CreateHistoryParams struct {
	IDAccount int64  `json:"id_account"`
	Amount    int64  `json:"amount"`
	Comment   string `json:"comment"`
}

func (q *Queries) CreateHistory(ctx context.Context, arg CreateHistoryParams) (History, error) {
	row := q.db.QueryRowContext(ctx, createHistory, arg.IDAccount, arg.Amount, arg.Comment)
	var i History
	err := row.Scan(
		&i.ID,
		&i.IDAccount,
		&i.Amount,
		&i.Comment,
		&i.CreatedAt,
	)
	return i, err
}

const deleteHistory = `-- name: DeleteHistory :exec
DELETE FROM history WHERE id = $1
`

func (q *Queries) DeleteHistory(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteHistory, id)
	return err
}

const getHistoryByAccountIdAmountSort = `-- name: GetHistoryByAccountIdAmountSort :many
SELECT id, id_account, amount, comment, created_at FROM history
WHERE id_account = $1
ORDER BY amount
LIMIT $2
OFFSET $3
`

type GetHistoryByAccountIdAmountSortParams struct {
	IDAccount int64 `json:"id_account"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

func (q *Queries) GetHistoryByAccountIdAmountSort(ctx context.Context, arg GetHistoryByAccountIdAmountSortParams) ([]History, error) {
	rows, err := q.db.QueryContext(ctx, getHistoryByAccountIdAmountSort, arg.IDAccount, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []History
	for rows.Next() {
		var i History
		if err := rows.Scan(
			&i.ID,
			&i.IDAccount,
			&i.Amount,
			&i.Comment,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getHistoryByAccountIdDataSort = `-- name: GetHistoryByAccountIdDataSort :many
SELECT id, id_account, amount, comment, created_at FROM history
WHERE id_account = $1
ORDER BY created_at
LIMIT $2
OFFSET $3
`

type GetHistoryByAccountIdDataSortParams struct {
	IDAccount int64 `json:"id_account"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

func (q *Queries) GetHistoryByAccountIdDataSort(ctx context.Context, arg GetHistoryByAccountIdDataSortParams) ([]History, error) {
	rows, err := q.db.QueryContext(ctx, getHistoryByAccountIdDataSort, arg.IDAccount, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []History
	for rows.Next() {
		var i History
		if err := rows.Scan(
			&i.ID,
			&i.IDAccount,
			&i.Amount,
			&i.Comment,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getHistoryById = `-- name: GetHistoryById :one
SELECT id, id_account, amount, comment, created_at FROM history
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetHistoryById(ctx context.Context, id int64) (History, error) {
	row := q.db.QueryRowContext(ctx, getHistoryById, id)
	var i History
	err := row.Scan(
		&i.ID,
		&i.IDAccount,
		&i.Amount,
		&i.Comment,
		&i.CreatedAt,
	)
	return i, err
}
