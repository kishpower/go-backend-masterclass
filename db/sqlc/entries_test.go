package db

import (
	"context"
	"testing"
	"time"

	"github.com/kishpower/simplebank/utils"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, acc Account) Entry {
	args := CreateEntryParams{
		AccountID: acc.ID,
		Amount:    utils.RandomMoney(),
	}

	entry1, err := testQueries.CreateEntry(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, entry1)

	require.Equal(t, args.AccountID, entry1.AccountID)
	require.Equal(t, args.Amount, entry1.Amount)

	require.NotZero(t, entry1.ID)
	require.NotZero(t, entry1.CreatedAt)

	return entry1
}

func TestCreateEntry(t *testing.T) {
	acc := CreateRandomAccount(t)
	createRandomEntry(t, acc)
}

func TestGetEntry(t *testing.T) {
	acc := CreateRandomAccount(t)
	entry1 := createRandomEntry(t, acc)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.Equal(t, entry1.ID, entry2.ID)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	acc := CreateRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomEntry(t, acc)
	}

	args := ListEntriesParams{
		AccountID: acc.ID,
		Limit:     5,
		Offset:    5,
	}
	entries, err := testQueries.ListEntries(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, args.AccountID, entry.AccountID)
	}
}
