package user_handler

type loginWithPhoneNumberRequest struct {
	PhoneNumber string `json:"phone_number"`
}

type verifyOTSSmsRequest struct {
	PhoneNumber string `json:"phone_number"`
	Code        string `json:"code"`
}
