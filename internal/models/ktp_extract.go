package models

type KTPExtract struct {
	Image           string `json:"image"`
	ValidateQuality bool   `json:"validate_quality"`
	ReferenceID     string `json:"reference_id"`
	CakkbackUrl     string `json:"callback_url"`
}
