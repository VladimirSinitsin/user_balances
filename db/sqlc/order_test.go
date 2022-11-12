package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/VladimirSinitsin/user_balances/util"
	"github.com/stretchr/testify/require"
)

func createRandomOrder(t *testing.T) Order {
	account := createRandomAccount(t)
	service, err := testQueries.GetServiceById(context.Background(), 1)
	if err != nil {
		service = createRandomService(t)
	}

	arg := CreateOrderParams{
		IDAccount:    account.ID,
		IDService:    service.ID,
		PriceService: service.Price,
		Status:       util.RandomStatus(),
	}

	order, err := testQueries.CreateOrder(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, order)

	require.Equal(t, arg.IDAccount, order.IDAccount)
	require.Equal(t, arg.IDService, order.IDService)
	require.Equal(t, arg.PriceService, order.PriceService)
	require.Equal(t, arg.Status, order.Status)

	require.NotZero(t, order.ID)
	require.NotZero(t, order.CreatedAt)

	return order
}

func TestQueries_CreateOrder(t *testing.T) {
	createRandomOrder(t)
}

func TestQueries_GetOrderById(t *testing.T) {
	order1 := createRandomOrder(t)
	order2, err := testQueries.GetOrderById(context.Background(), order1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, order2)

	require.Equal(t, order1.ID, order2.ID)
	require.Equal(t, order1.IDAccount, order2.IDAccount)
	require.Equal(t, order1.IDService, order2.IDService)
	require.Equal(t, order1.PriceService, order2.PriceService)
	require.Equal(t, order1.Status, order2.Status)
	require.WithinDuration(t, order1.CreatedAt, order2.CreatedAt, time.Second)
}

func TestQueries_UpdateOrder(t *testing.T) {
	order1 := createRandomOrder(t)

	arg := UpdateOrderParams{
		ID:     order1.ID,
		Status: util.RandomStatus(),
	}

	order2, err := testQueries.UpdateOrder(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, order2)

	require.Equal(t, order1.ID, order2.ID)
	require.Equal(t, order1.IDAccount, order2.IDAccount)
	require.Equal(t, order1.IDService, order2.IDService)
	require.Equal(t, order1.PriceService, order2.PriceService)
	require.Equal(t, arg.Status, order2.Status)
	require.WithinDuration(t, order1.CreatedAt, order2.CreatedAt, time.Second)
}

func TestQueries_DeleteOrder(t *testing.T) {
	order1 := createRandomOrder(t)
	err := testQueries.DeleteOrder(context.Background(), order1.ID)
	require.NoError(t, err)

	order2, err := testQueries.GetOrderById(context.Background(), order1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, order2)
}

func TestQueries_ListOrdersByServiceId(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomOrder(t)
	}

	idService := int64(1)
	orders, err := testQueries.ListOrdersByServiceId(context.Background(), idService)
	require.NoError(t, err)

	for _, order := range orders {
		require.NotEmpty(t, order)
	}
}
