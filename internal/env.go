package internal

import (
	"fmt"
	"os"

	"github.com/lpernett/godotenv"
)

var envVars map[string]string

func InitEnvVars() map[string]string {
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

	supabaseKey := os.Getenv("SUPABASE_KEY")
	if supabaseKey == "" {
		panic("Error: SUPABASE_KEY not found in .env")
	}

	supabaseURL := os.Getenv("SUPABASE_URL")
	if supabaseURL == "" {
		panic("Error: SUPABASE_URL not found in .env")
	}

	myPhoneNumber := os.Getenv("MY_PHONE_NUMBER")
	if myPhoneNumber == "" {
		panic("Error: MY_PHONE_NUMBER not found in .env")
	}

	twilioPhoneNumber := os.Getenv("TWILIO_PHONE_NUMBER")
	if twilioPhoneNumber == "" {
		panic("Error: TWILIO_PHONE_NUMBER not found in .env")
	}

	envVars["YNAB_KEY"] = ynabSecret
	envVars["BUDGET_ID"] = budgetID
	envVars["TWILIO_ACCOUNT_SID"] = twilioAccountSid
	envVars["TWILIO_AUTH_TOKEN"] = twilioAuthToken
	envVars["SUPABASE_KEY"] = supabaseKey
	envVars["SUPABASE_URL"] = supabaseURL
	envVars["MY_PHONE_NUMBER"] = myPhoneNumber
	envVars["TWILIO_PHONE_NUMBER"] = twilioPhoneNumber

	return envVars
}
