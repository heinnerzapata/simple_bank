package simplebank

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
