package models

type VerifyFace struct {
	Image             string `json:"image"`
	SubjectID         string `json:"subject_id"`
	IsQuality         bool   `json:"is_quality"`
	IsAttribute       bool   `json:"is_attribute"`
	IsLiveness        bool   `json:"is_liveness"`
	ValidateQuality   bool   `json:"validate_quality"`
	ValidateAttribute bool   `json:"validate_attribute"`
	ValidateLiveness  bool   `json:"validate_liveness"`
	Threshold         string `json:"threshhold"`
	ValidateNface     bool   `json:"validate_nface"`
}

type Liveness struct {
	Probability string `json:"probability"`
	Status      string `json:"status"`
}
