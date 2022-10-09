package twilio_config

import (
	"os"

	"github.com/twilio/twilio-go"
)

type TwilioClient struct {
	Client     *twilio.RestClient
	ServiceSID string
}

func NewTwilioClient() *TwilioClient {
	return &TwilioClient{
		Client:     twilio.NewRestClient(),
		ServiceSID: os.Getenv("VERIFY_SERVICE_SID"),
	}
}
