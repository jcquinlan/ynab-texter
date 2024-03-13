package messaging

import (
	"fmt"
	"jcquinlan/ynab-texter/env"

	"github.com/brunomvsouza/ynab.go/api/transaction"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func sendText(body string) {
	envVars := env.GetEnvVars()

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: envVars["TWILIO_ACCOUNT_SID"],
		Password: envVars["TWIlIO_AUTH_TOKEN"],
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo("+12678846019")
	params.SetFrom("+12674583734")
	params.SetBody(body)

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println("Error sending SMS message: " + err.Error())
	}
}

func SendTransactionMessage(transaction *transaction.Transaction) {
	cost := "$" + fmt.Sprint(transaction.Amount/1000)
	body := "Charge for " + cost + " from " + *transaction.PayeeName
	sendText(body)
}
