package models

type Face struct {
	Image             string  `json:"image"`
	SubjectID         *string `json:"subject_id"`
	IsQuality         bool    `json:"is_quality"`
	IsAttribute       bool    `json:"is_attribute"`
	IsLiveness        bool    `json:"is_liveness"`
	ValidateQuality   bool    `json:"validate_quality"`
	ValidateAttribute bool    `json:"validate_attribute"`
	ValidateLiveness  bool    `json:"validate_liveness"`
	ValidateNface     bool    `json:"validate_nface"`
}

type CompareFace struct {
	Image1            string `json:"image1"`
	Image2            string `json:"image2"`
	IsQuality         bool   `json:"is_quality"`
	IsAttribute       bool   `json:"is_attribute"`
	IsLiveness        bool   `json:"is_liveness"`
	ValidateQuality   bool   `json:"validate_quality"`
	ValidateAttribute bool   `json:"validate_attribute"`
	ValidateLiveness  bool   `json:"validate_liveness"`
	ValidateNface     bool   `json:"validate_nface"`
	Threshold         string `json:"threshold"`
}

type BoundingBox struct {
	TopLeftX     string
	TopLeftY     string
	BottomRightX string
	BottomRightY string
	Width        string
	Height       string
}

type FaceLandmark struct {
	LeftEyeX    string
	LeftEyeY    string
	RightEyeX   string
	RightEyeY   string
	NoseX       string
	NoseY       string
	MouthLeftX  string
	MouthLeftY  string
	MouthRightX string
	MouthRightY string
}

type FaceAttribute struct {
	SunglassesOn bool `json:"sunglasses_on"`
	VeilOn       bool `json:"veil_on"`
	MaskOn       bool `json:"mask_on"`
	HatOn        bool `json:"haton"`
}

type ImageQuality struct {
	Blur      bool `json:"blur"`
	Dark      bool `json:"dark"`
	Grayscale bool `json:"grayscale"`
}

type FaceData struct {
	Message      string        `json:"message"`
	SessionID    string        `json:"session_id"`
	Timestamp    int64         `json:"timestamp"`
	StatusCode   string        `json:"status_code"`
	BoundingBox  BoundingBox   `json:"bounding_box"`
	FaceLandmark FaceLandmark  `json:"face_landmark"`
	Rotation     int           `json:"rotation"`
	Nface        int           `json:"nface"`
	Attibutes    FaceAttribute `json:"attributes"`
	ImageQuality ImageQuality  `json:"image_quality"`
	SubjectID    string        `json:"subject_id"`
}

type FaceDataResponse struct {
	Message string   `json:"message"`
	Data    FaceData `json:"data"`
}

type FaceDataDetect struct {
	Attibutes    FaceAttribute `json:"attributes"`
	BoundingBox  BoundingBox   `json:"bounding_box"`
	FaceLandmark FaceLandmark  `json:"face_landmark"`
	SessionID    string        `json:"session_id"`
	Timestamp    string        `json:"timestamp"`
	ImageQuality ImageQuality  `json:"image_quality"`
	Rotation     int           `json:"rotation"`
	Nface        int           `json:"nface"`
	SubjectID    string        `json:"subject_id"`
}

type FaceDataDetectResponse struct {
	StatusCode string         `json:"status_code"`
	Response   FaceDataDetect `json:"response"`
}

type FaceDetectResponse struct {
	Message string                 `json:"message"`
	Data    FaceDataDetectResponse `json:"data"`
}
