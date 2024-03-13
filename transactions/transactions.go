package transactions

import (
	"jcquinlan/ynab-texter/env"
	"time"

	"github.com/brunomvsouza/ynab.go"
	"github.com/brunomvsouza/ynab.go/api"
	"github.com/brunomvsouza/ynab.go/api/transaction"
)

func GetRecentTransactions() ([]*transaction.Transaction, error) {
	envVars := env.GetEnvVars()

	ynabClient := ynab.NewClient(envVars["YNAB_KEY"])
	filterTime := time.Now().
		Add(-time.Hour * 24 * 1)

	transactionFilter := transaction.Filter{
		Since: &api.Date{Time: filterTime},
		Type:  transaction.StatusUnapproved.Pointer(),
	}
	transactions, err := ynabClient.Transaction().GetTransactions(
		envVars["BUDGET_ID"],
		&transactionFilter,
	)

	if err != nil {
		return nil, err
	}

	return transactions, nil
}
