package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store providers all functions to execute db queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

// TransferTxParams contains the input parameters of the transaction
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// TransferTxResults contains the result parameters of the transaction
type TransferTxResults struct {
	HistoryAccountFrom History `json:"history_account_from"`
	HistoryAccountTo   History `json:"history_account_to"`
	FromAccount        Account `json:"from_account"`
	ToAccount          Account `json:"to_account"`
	Amount             int64   `json:"amount"`
	OperationStatus    string  `json:"operation_status"`
	Date               string  `json:"date"`
}

// TransferTx performs a money transfer from one account to the other.
// It creates a records in the History table and update accounts' balance with a single database transaction.
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResults, error) {
	var result TransferTxResults

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.HistoryAccountFrom, err = q.CreateHistory(ctx, CreateHistoryParams{
			IDAccount: arg.FromAccountID,
			Amount:    -arg.Amount,
			Comment:   fmt.Sprintf("Перевод клиенту № %v", arg.ToAccountID),
		})
		if err != nil {
			return err
		}

		result.HistoryAccountTo, err = q.CreateHistory(ctx, CreateHistoryParams{
			IDAccount: arg.ToAccountID,
			Amount:    arg.Amount,
			Comment:   fmt.Sprintf("Перевод от клиента № %v", arg.FromAccountID),
		})
		if err != nil {
			return err
		}

		// update accounts' balance
		result.FromAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			ID:     arg.FromAccountID,
			Amount: -arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ToAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			ID:     arg.ToAccountID,
			Amount: arg.Amount,
		})
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
