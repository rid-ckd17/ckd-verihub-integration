package models

type ResponseSendOtp struct {
	Message      string `json:"message"`
	MSISDN       string `json:"msisdn"`
	OTP          string `json:"otp"`
	SegmentCount int    `json:"segment_count"`
	SessionID    string `json:"session_id"`
}
