package main

import (
	"fmt"
	"time"

	"github.com/brunomvsouza/ynab.go"
	"github.com/brunomvsouza/ynab.go/api"
	"github.com/brunomvsouza/ynab.go/api/transaction"
)

type ynabClient struct {
	client ynab.ClientServicer
}

func CreateYnabClient() *ynabClient {
	return &ynabClient{
		client: ynab.NewClient(envVars["YNAB_KEY"]),
	}
}

func (y *ynabClient) GetRecentTransactions() ([]*transaction.Transaction, error) {
	fmt.Println("Getting recent transactions from YNAB...")
	filterTime := time.Now().
		Add(-time.Hour * 24 * 2)

	transactionFilter := transaction.Filter{
		Since: &api.Date{Time: filterTime},
		Type:  transaction.StatusUnapproved.Pointer(),
	}
	transactions, err := y.client.Transaction().GetTransactions(
		envVars["BUDGET_ID"],
		&transactionFilter,
	)

	if err != nil {
		return nil, err
	}

	return transactions, nil
}
