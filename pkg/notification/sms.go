package notification

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendMessage(phoneNumber string, message string) {
	sender := viper.GetString("TWILIO_PHONE_NUMBER")
	accountSID := viper.GetString("TWILIO_ACCOUNT_SID")
	authToken := viper.GetString("TWILIO_AUTH_TOKEN")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{Username: accountSID, Password: authToken})
	params := &twilioApi.CreateMessageParams{}
	params.SetTo(phoneNumber)
	params.SetFrom(sender)
	params.SetBody(message)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println("Error sending SMS message: ")
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Printf("Response: %s", string(response))
	}
}
