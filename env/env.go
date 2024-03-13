package env

import (
	"fmt"
	"os"

	"github.com/lpernett/godotenv"
)

var envVars map[string]string

func GetEnvVars() map[string]string {

	if len(envVars) > 0 {
		return envVars
	}

	envVars = make(map[string]string)

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	ynabSecret := os.Getenv("YNAB_KEY")
	if ynabSecret == "" {
		panic("Error: YNAB_KEY not found in .env")
	}

	budgetID := os.Getenv("BUDGET_ID")
	if budgetID == "" {
		panic("Error: BUDGET_ID not found in .env")
	}

	twilioAccountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	if twilioAccountSid == "" {
		panic("Error: TWILIO_ACCOUNT_SID not found in .env")
	}

	twilioAuthToken := os.Getenv("TWILIO_AUTH_TOKEN")
	if twilioAuthToken == "" {
		panic("Error: TWILIO_AUTH_TOKEN not found in .env")
	}

	envVars["YNAB_KEY"] = ynabSecret
	envVars["BUDGET_ID"] = budgetID
	envVars["TWILIO_ACCOUNT_SID"] = twilioAccountSid
	envVars["TWILIO_AUTH_TOKEN"] = twilioAuthToken

	return envVars
}
