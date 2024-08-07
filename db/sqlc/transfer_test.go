package db

import (
	"context"
	"testing"
	"time"

	"github.com/fzl-22/simple-bank/utils"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, srcAccount, dstAccount Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: srcAccount.ID,
		ToAccountID:   dstAccount.ID,
		Amount:        utils.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	srcAccount := createRandomAccount(t)
	dstAccount := createRandomAccount(t)

	createRandomTransfer(t, srcAccount, dstAccount)
}

func TestGetTransfer(t *testing.T) {
	srcAccount := createRandomAccount(t)
	dstAccount := createRandomAccount(t)

	transfer1 := createRandomTransfer(t, srcAccount, dstAccount)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	srcAccount := createRandomAccount(t)
	dstAccount := createRandomAccount(t)

	var transfers1 []Transfer
	for i := 0; i < 5; i++ {
		transfer := createRandomTransfer(t, srcAccount, dstAccount)
		transfers1 = append(transfers1, transfer)
	}

	arg := ListTransfersParams{
		FromAccountID: srcAccount.ID,
		ToAccountID:   dstAccount.ID,
		Limit:         5,
		Offset:        0,
	}

	transfers2, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfers2)

	for index, transfer := range transfers2 {
		require.Equal(t, transfers1[index].ID, transfer.ID)
		require.Equal(t, transfers1[index].FromAccountID, transfer.FromAccountID)
		require.Equal(t, transfers1[index].ToAccountID, transfer.ToAccountID)
		require.Equal(t, transfers1[index].Amount, transfer.Amount)
		require.WithinDuration(t, transfers1[index].CreatedAt, transfer.CreatedAt, time.Second)
	}
}
