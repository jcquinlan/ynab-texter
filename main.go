package main

import (
	"fmt"

	ynabSDK "github.com/brunomvsouza/ynab.go/api/transaction"
)

func main() {
	InitEnvVars()
	supabase := CreateSupabaseClient()
	ynab := CreateYnabClient()
	messaging := InitMessagingClient()

	recentTransactions, err := ynab.GetRecentTransactions()
	if err != nil {
		panic(err)
	}

	if len(recentTransactions) == 0 {
		fmt.Println("No recent transactions - exiting.")
		return
	}

	recentTransactionIds := make([]string, len(recentTransactions))
	for i, transaction := range recentTransactions {
		recentTransactionIds[i] = transaction.ID
	}

	transactionRecordsForRecentTransactions, err := supabase.GetTransactionRecords(recentTransactionIds)
	if err != nil {
		panic(err)
	}

	newRecentTransactions := make([]*ynabSDK.Transaction, 0)

	fmt.Println("Calculating new transactions...")
	for _, transaction := range recentTransactions {
		found := false
		for _, record := range transactionRecordsForRecentTransactions {
			if transaction.ID == record.TransactionId {
				found = true
				break
			}
		}
		if !found {
			newRecentTransactions = append(newRecentTransactions, transaction)
		}
	}

	if len(newRecentTransactions) == 0 {
		fmt.Println("No new transactions - exiting.")
		messaging.SendText("No new transactions today")
		return
	}

	for _, transaction := range newRecentTransactions {
		newTransactionRecord := TransactionRecordDTO{
			TransactionId: transaction.ID,
		}

		insertedTransactionRecords, err := supabase.CreateTransactionRecord(&newTransactionRecord)
		if err != nil {
			panic(err)
		}

		if len(insertedTransactionRecords) > 0 {
			messaging.SendTransactionMessage(transaction)
		}
	}
}
