package main

import (
	"fmt"

	"github.com/brunomvsouza/ynab.go/api/transaction"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type Messaging struct {
	client *twilio.RestClient
}

func InitMessagingClient() *Messaging {
	messagingClient := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: envVars["TWILIO_ACCOUNT_SID"],
		Password: envVars["TWILIO_AUTH_TOKEN"],
	})

	return &Messaging{
		client: messagingClient,
	}
}

func (m *Messaging) SendText(body string) {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: envVars["TWILIO_ACCOUNT_SID"],
		Password: envVars["TWIlIO_AUTH_TOKEN"],
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(envVars["MY_PHONE_NUMBER"])
	params.SetFrom(envVars["TWILIO_PHONE_NUMBER"])
	params.SetBody(body)

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println("Error sending SMS message: " + err.Error())
	}
}

func (m *Messaging) SendTransactionMessage(transaction *transaction.Transaction) {
	fmt.Println("Sending transaction message for transaction: ", transaction.ID)
	cost := "$" + fmt.Sprint(transaction.Amount/1000)
	body := "Charge for " + cost + " from " + *transaction.PayeeName
	m.SendText(body)
}
