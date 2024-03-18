package main

import (
	"fmt"

	supa "github.com/nedpals/supabase-go"
)

type Supabase struct {
	client *supa.Client
}

type TransactionRecordDTO struct {
	TransactionId string `json:"transaction_id"`
}

type TransactionRecord struct {
	ID        int32  `json:"id"`
	CreatedAt string `json:"created_at"`
	Status    string `json:"status"`
	TransactionRecordDTO
}

func (t *TransactionRecord) ReadableName() string {
	return fmt.Sprintf("TransactionRecord %d", t.ID)
}

func CreateSupabaseClient() *Supabase {
	var supabaseKey = envVars["SUPABASE_KEY"]
	var supabaseURL = envVars["SUPABASE_URL"]

	var SupabaseClient = supa.CreateClient(supabaseURL, supabaseKey)
	return &Supabase{client: SupabaseClient}
}

func (s *Supabase) GetTransactionRecords(transactionRecordIds []string) ([]*TransactionRecord, error) {
	fmt.Println("Getting transaction records for ids: ", transactionRecordIds)
	var records []*TransactionRecord
	err := s.client.DB.From("transaction_records").Select("*").In("transaction_id", transactionRecordIds).Execute(&records)

	return records, err
}

func (s *Supabase) CreateTransactionRecord(transactionRecordDTO *TransactionRecordDTO) ([]*TransactionRecord, error) {
	fmt.Println("Creating transaction record for transaction id: ", transactionRecordDTO.TransactionId)
	var insertedRecords []*TransactionRecord
	err := s.client.DB.From("transaction_records").Insert(transactionRecordDTO).Execute(&insertedRecords)

	return insertedRecords, err
}
