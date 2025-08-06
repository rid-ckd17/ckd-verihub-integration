package models

type VerificationData struct {
	NIK         string `json:"nik"`
	Name        string `json:"name"`
	BirthDate   string `json:"birth_date"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	SelfiePhoto string `json:"selfie_photo"`
	KtpPhoto    string `json:"ktp_photo"`
	Channel     string `json:"channel"`
	ReferenceID string `json:"reference_id"`
}

type ResponseSucessVefification struct {
	Message string       `json:"message"`
	Data    Verification `json:"data"`
}

type Verification struct {
	ID          string   `json:"id"`
	Status      string   `json:"status"`
	RejectField []string `json:"reject_field"`
	ReferenceID string   `json:"reference_id"`
}

type ResponseInvalid struct {
	Status     int          `json:"status"`
	Message    string       `json:"message"`
	ErrorCode  string       `json:"error_code"`
	ErrorField []ErrorField `json:"error_field"`
}

type ErrorField struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
