package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	args := CreateAccountParams{
		Balance:  100,
		Owner:    "John",
		Currency: "USD",
	}

	account, err := testQueries.CreateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, args.Balance, account.Balance)
	require.Equal(t, args.Currency, account.Currency)
	require.Equal(t, args.Owner, account.Owner)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}
