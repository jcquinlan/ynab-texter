package main

import (
	"jcquinlan/ynab-texter/messaging"
	"jcquinlan/ynab-texter/transactions"
)

func main() {
	recentTransactions, err := transactions.GetRecentTransactions()

	if err != nil {
		panic(err)
	}

	for _, transaction := range recentTransactions {
		messaging.SendTransactionMessage(transaction)
	}
}
