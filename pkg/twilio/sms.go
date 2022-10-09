package twilio

import (
	"andalas-care/configs/twilio_config"

	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

func SendOTPSms(phoneNumber string, myTwilio *twilio_config.TwilioClient) (status string, err error) {

	client := myTwilio.Client

	var (
		channel = "sms"
	)

	params := &openapi.CreateVerificationParams{
		To:      &phoneNumber,
		Channel: &channel,
	}

	resp, err := client.VerifyV2.CreateVerification(myTwilio.ServiceSID, params)

	return *resp.Status, err
}

func VerifyOTPSms(phoneNumber, code string, myTwilio *twilio_config.TwilioClient) (status bool, err error) {

	client := myTwilio.Client

	params := &openapi.CreateVerificationCheckParams{
		Code: &code,
		To:   &phoneNumber,
	}

	resp, err := client.VerifyV2.CreateVerificationCheck(myTwilio.ServiceSID, params)

	return *resp.Status == "approved", err
}
