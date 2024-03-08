package db

import (
	"bank/util"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func CreateRandomPersonal(t *testing.T, account Account) Personal{
	arg := CreatePersonalParams{
		AccountID: account.ID,
		Amount: util.RandomMoney(),
	}

	personal, err := testQueries.CreatePersonal(context.Background(),arg)
	require.NoError(t,err)
	require.NotEmpty(t,personal)

	require.Equal(t, arg.AccountID, personal.AccountID)
	require.Equal(t, arg.Amount, personal.Amount)

	require.NotZero(t,personal.ID)
	require.NotZero(t,personal.CreatedAt)

	return personal
	
}

func TestCreatePersonal(t *testing.T){
	Account := createRandomAccount(t)
	CreateRandomPersonal(t, Account)
}

func TestGetPersonal(t *testing.T){
	Account := createRandomAccount(t)
	personal1 := CreateRandomPersonal(t, Account) 
	personal2 , err := testQueries.GetPersonal(context.Background(),personal1.ID)
	require.NoError(t,err )
	require.NotEmpty(t,personal2)

	require.Equal(t, personal1.AccountID, personal2.AccountID)
	require.Equal(t, personal1.Amount, personal2.Amount)
	require.WithinDuration(t, personal1.CreatedAt,personal2.CreatedAt, time.Second)
}

func TestListpersonal(t *testing.T){
	Account := createRandomAccount(t)
	for i := 0; i < 10 ; i ++ {
		CreateRandomPersonal(t,Account)
	}

	arg := ListpersonalParams {
		Limit: 5,
		Offset: 5,
	}

	Personal, err := testQueries.Listpersonal(context.Background(),arg)
	require.NoError(t,err)
	require.Len(t, Personal, 5)

	for _, personal := range Personal {
		require.NotEmpty(t,personal)
	}
}