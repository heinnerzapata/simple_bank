package db

import (
	"context"
	"database/sql"
	"simplebank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RanddomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func createRandomEntrie(t *testing.T) Entry {

	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	arg := CreateEntrieParams{
		AccountID: account2.ID,
		Amount:    util.RandomMoney(),
	}

	entrie, err := testQueries.CreateEntrie(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entrie)

	require.Equal(t, arg.AccountID, entrie.AccountID)
	require.Equal(t, arg.Amount, entrie.Amount)

	require.NotZero(t, entrie.AccountID)
	require.NotZero(t, entrie.Amount)

	return entrie
}

func createRandomTransfer(t *testing.T) Transfer {

	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	account3 := createRandomAccount(t)
	account4, err := testQueries.GetAccount(context.Background(), account3.ID)

	require.NoError(t, err)

	arg := CreateTransferParams{
		FromAccountID: account2.ID,
		ToAccountID:   account4.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.FromAccountID)
	require.NotZero(t, transfer.ToAccountID)
	require.NotZero(t, transfer.Amount)

	return transfer
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: util.RandomMoney(),
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, arg.Limit, int32(len(accounts)))

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}

// Test Entries

func TestGetEntrie(t *testing.T) {
	entrie1 := createRandomEntrie(t)

	entrie2, err := testQueries.GetEntrie(context.Background(), entrie1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entrie2)
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntrie(t)
	}

	arg := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, arg.Limit, int32(len(entries)))

	for _, entrie := range entries {
		require.NotEmpty(t, entrie)
	}
}

func TestUpdateEntrie(t *testing.T) {
	entrie1 := createRandomEntrie(t)

	arg := UpdateEntrieParams{
		ID:     entrie1.ID,
		Amount: util.RandomMoney(),
	}

	entrie2, err := testQueries.UpdateEntrie(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entrie2)

	require.Equal(t, entrie1.ID, entrie2.ID)
	require.Equal(t, arg.Amount, entrie2.Amount)
}

func TestDeleteEntrie(t *testing.T) {
	entrie1 := createRandomEntrie(t)
	err := testQueries.DeleteAccount(context.Background(), entrie1.ID)
	require.NoError(t, err)

	entrie2, err := testQueries.GetAccount(context.Background(), entrie1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, entrie2)
}

// Test Transfers

func TestGetTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)
}

func TestListTransfers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTransfer(t)
	}

	arg := ListTransfersParams{
		Limit:  5,
		Offset: 5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, arg.Limit, int32(len(transfers)))

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}

func TestUpdateTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)

	arg := UpdateTransferParams{
		ID:     transfer1.ID,
		Amount: util.RandomMoney(),
	}

	transfer2, err := testQueries.UpdateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, arg.Amount, transfer2.Amount)
}

func TestDeleteTransfer(t *testing.T) {
	entrie1 := createRandomEntrie(t)
	err := testQueries.DeleteAccount(context.Background(), entrie1.ID)
	require.NoError(t, err)

	entrie2, err := testQueries.GetAccount(context.Background(), entrie1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, entrie2)
}
