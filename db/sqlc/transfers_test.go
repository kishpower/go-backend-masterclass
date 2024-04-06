package db

import (
	"context"
	"testing"
	"time"

	"github.com/kishpower/simplebank/utils"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, acc1 Account, acc2 Account) Transfer {
	args := CreateTransferParams{
		FromAccountID: acc1.ID,
		ToAccountID:   acc2.ID,
		Amount:        utils.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, args.FromAccountID, transfer.FromAccountID)
	require.Equal(t, args.ToAccountID, transfer.ToAccountID)
	require.Equal(t, args.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	acc1 := CreateRandomAccount(t)
	acc2 := CreateRandomAccount(t)

	createRandomTransfer(t, acc1, acc2)
}

func TestGetTransfer(t *testing.T) {
	acc1 := CreateRandomAccount(t)
	acc2 := CreateRandomAccount(t)
	transfer1 := createRandomTransfer(t, acc1, acc2)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	acc1 := CreateRandomAccount(t)
	acc2 := CreateRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomTransfer(t, acc1, acc2)
	}

	args := ListTransfersParams{
		FromAccountID: acc1.ID,
		ToAccountID:   acc2.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}
