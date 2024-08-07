package db

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/fzl-22/simple-bank/utils"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, account Account) Entry {
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    utils.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntry(t, account)
}

func TestGetEntry(t *testing.T) {
	account := createRandomAccount(t)
	entry1 := createRandomEntry(t, account)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	account := createRandomAccount(t)

	var entries1 []Entry
	for i := 0; i < 10; i++ {
		entry := createRandomEntry(t, account)
		entries1 = append(entries1, entry)
	}

	arg := ListEntriesParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    0,
	}

	entries2, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	fmt.Println(len(entries2))
	require.Len(t, entries2, 5)

	for index, entry := range entries2 {
		require.Equal(t, entries1[index].ID, entry.ID)
		require.Equal(t, entries1[index].AccountID, entry.AccountID)
		require.Equal(t, entries1[index].Amount, entry.Amount)
		require.WithinDuration(t, entries1[index].CreatedAt, entry.CreatedAt, time.Second)
	}
}
