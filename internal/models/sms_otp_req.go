package models

type RequestSmsOtp struct {
	MSISDN      string  `json:"msisdn"`
	Template    *string `json:"template"`
	OTP         *string `json:"otp"`
	TimeLimit   *string `json:"time_limit"`
	Challenge   *string `json:"challenge"`
	CallbackURL *string `json:"callback_url"`
}

type RequestVerifySmsOtp struct {
	MSISDN    string `json:"msisdn"`
	OTP       string `json:"otp"`
	Challenge string `json:"challenge"`
}
