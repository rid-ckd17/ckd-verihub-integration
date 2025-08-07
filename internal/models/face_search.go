package models

type SearchFace struct {
	Image             string `json:"image"`
	Limit             int    `json:"limit"`
	ReturnImage       bool   `json:"return_image"`
	Threshhold        int    `json:"threshold"`
	IsQuality         bool   `json:"is_quality"`
	IsAttribute       bool   `json:"is_attribute"`
	IsLiveness        bool   `json:"is_liveness"`
	ValidateQuality   bool   `json:"validate_quality"`
	ValidateAttribute bool   `json:"validate_attribute"`
	ValidateLiveness  bool   `json:"validate_liveness"`
	ValidateNface     bool   `json:"validate_nface"`
}
