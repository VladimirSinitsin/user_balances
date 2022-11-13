package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/VladimirSinitsin/user_balances/util"
	"github.com/stretchr/testify/require"
)

func createRandomHistory(t *testing.T) History {
	account, err := testQueries.GetAccount(context.Background(), 1)
	if err != nil {
		account = createRandomAccount(t)
	}

	arg := CreateHistoryParams{
		IDAccount: account.ID, // 1
		Amount:    util.RandomAmount(),
		Comment:   util.RandomComment(),
	}

	history, err := testQueries.CreateHistory(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, history)

	require.Equal(t, arg.IDAccount, history.IDAccount)
	require.Equal(t, arg.Amount, history.Amount)
	require.Equal(t, arg.Comment, history.Comment)

	require.NotZero(t, history.ID)
	require.NotZero(t, history.CreatedAt)

	return history
}

func TestQueries_CreateHistory(t *testing.T) {
	createRandomHistory(t)
}

func TestQueries_GetHistoryById(t *testing.T) {
	history1 := createRandomHistory(t)
	history2, err := testQueries.GetHistoryById(context.Background(), history1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, history2)

	require.Equal(t, history1.ID, history2.ID)
	require.Equal(t, history1.IDAccount, history2.IDAccount)
	require.Equal(t, history1.Amount, history2.Amount)
	require.Equal(t, history1.Comment, history2.Comment)
	require.WithinDuration(t, history1.CreatedAt, history2.CreatedAt, time.Second)
}

func TestQueries_DeleteHistory(t *testing.T) {
	history1 := createRandomHistory(t)
	err := testQueries.DeleteHistory(context.Background(), history1.ID)
	require.NoError(t, err)

	history2, err := testQueries.GetHistoryById(context.Background(), history1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, history2)
}

func TestQueries_ListHistoryByAccountIdAmountSort(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomHistory(t)
	}

	arg := ListHistoryByAccountIdAmountSortParams{
		IDAccount: int64(1),
		Limit:     5,
		Offset:    5,
	}

	histories, err := testQueries.ListHistoryByAccountIdAmountSort(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, histories, 5)

	for _, history := range histories {
		require.NotEmpty(t, history)
	}
}

func TestQueries_ListHistoryByAccountIdDataSort(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomHistory(t)
	}

	arg := ListHistoryByAccountIdDataSortParams{
		IDAccount: int64(1),
		Limit:     5,
		Offset:    5,
	}

	histories, err := testQueries.ListHistoryByAccountIdDataSort(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, histories, 5)

	for _, history := range histories {
		require.NotEmpty(t, history)
	}
}
