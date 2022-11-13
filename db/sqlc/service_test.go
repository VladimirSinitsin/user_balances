package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/VladimirSinitsin/user_balances/util"
	"github.com/stretchr/testify/require"
)

func createRandomService(t *testing.T) Service {
	arg := CreateServiceParams{
		Name:  util.RandomName(),
		Price: util.RandomAmount(),
	}

	service, err := testQueries.CreateService(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, service)

	require.Equal(t, arg.Name, service.Name)
	require.Equal(t, arg.Price, service.Price)

	require.NotZero(t, service.ID)
	require.NotZero(t, service.CreatedAt)

	return service
}

func TestQueries_CreateService(t *testing.T) {
	createRandomService(t)
}

func TestQueries_GetServiceById(t *testing.T) {
	service1 := createRandomService(t)
	service2, err := testQueries.GetServiceById(context.Background(), service1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, service2)

	require.Equal(t, service1.ID, service2.ID)
	require.Equal(t, service1.Name, service2.Name)
	require.Equal(t, service1.Price, service2.Price)
	require.WithinDuration(t, service1.CreatedAt, service2.CreatedAt, time.Second)
}

func TestQueries_GetServiceByName(t *testing.T) {
	service1 := createRandomService(t)
	service2, err := testQueries.GetServiceByName(context.Background(), service1.Name)
	require.NoError(t, err)
	require.NotEmpty(t, service2)

	require.Equal(t, service1.ID, service2.ID)
	require.Equal(t, service1.Name, service2.Name)
	require.Equal(t, service1.Price, service2.Price)
	require.WithinDuration(t, service1.CreatedAt, service2.CreatedAt, time.Second)
}

func TestQueries_DeleteService(t *testing.T) {
	service1 := createRandomService(t)
	err := testQueries.DeleteService(context.Background(), service1.ID)
	require.NoError(t, err)

	service2, err := testQueries.GetServiceById(context.Background(), service1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, service2)
}

func TestQueries_ListServices(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomService(t)
	}

	services, err := testQueries.ListServices(context.Background())
	require.NoError(t, err)

	for _, service := range services {
		require.NotEmpty(t, service)
	}
}
