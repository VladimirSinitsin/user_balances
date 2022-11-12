package db

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStore_TransferTx(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	// run n concurrent transfer transactions
	n := 5
	amount := int64(10)
	testCommentFrom := fmt.Sprintf("Перевод клиенту № %v", account2.ID)
	testCommentTo := fmt.Sprintf("Перевод от клиента № %v", account1.ID)

	errs := make(chan error)
	results := make(chan TransferTxResults)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})

			errs <- err
			results <- result
		}()
	}

	// check results
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// check history of `account1`
		historyAccountFrom := result.HistoryAccountFrom
		require.NotEmpty(t, historyAccountFrom)
		require.Equal(t, account1.ID, historyAccountFrom.IDAccount)
		require.Equal(t, -amount, historyAccountFrom.Amount)
		require.Equal(t, testCommentFrom, historyAccountFrom.Comment)
		require.NotZero(t, historyAccountFrom.ID)
		require.NotZero(t, historyAccountFrom.CreatedAt)

		_, err = store.GetHistoryById(context.Background(), historyAccountFrom.ID)
		require.NoError(t, err)

		// check history of `account2`
		historyAccountTo := result.HistoryAccountTo
		require.NotEmpty(t, historyAccountTo)
		require.Equal(t, account2.ID, historyAccountTo.IDAccount)
		require.Equal(t, amount, historyAccountTo.Amount)
		require.Equal(t, testCommentTo, historyAccountTo.Comment)
		require.NotZero(t, historyAccountTo.ID)
		require.NotZero(t, historyAccountTo.CreatedAt)

		_, err = store.GetHistoryById(context.Background(), historyAccountTo.ID)
		require.NoError(t, err)

		// TODO: check accounts' balance
	}
}
