package verihub

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testapiverihub/internal/models"
)

type VerihubSdk struct {
	AppID  string
	ApiKey string
	URLV1  string
	URLV2  string
	Ctx    *context.Context
	Mode   int
}

func NewVirehubSdk(appId string, apiKey string, ctx *context.Context, mode int, urlv1 string, urlv2 string) *VerihubSdk {
	return &VerihubSdk{AppID: appId, ApiKey: apiKey, Ctx: ctx, Mode: mode, URLV1: urlv1, URLV2: urlv2}
}

func (v *VerihubSdk) clientRequest(mode string, data string, path string, version string) ([]byte, *int, error) {

	url := ""

	if version == "v1" {
		url = v.URLV1 + path + mode
	} else if version == "v2" {
		url = v.URLV2 + path + mode
	}

	payload := strings.NewReader(data)
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("App-ID", v.AppID)
	req.Header.Add("API-Key", v.ApiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, &res.StatusCode, err
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))
	return body, &res.StatusCode, nil
}

func (v *VerihubSdk) EnrollFace(face models.Face) (*models.FaceDataResponse, *int, error) {
	mode := ""

	if v.Mode == 0 {
		mode = ""
	}

	parse, err := json.Marshal(&face)
	if err != nil {
		return nil, nil, err
	}

	body, statusCode, err := v.clientRequest(mode, string(parse), "/face/enroll", "v1")
	if err != nil {
		return nil, statusCode, err
	}

	var data models.FaceDataResponse

	if err := json.Unmarshal(body, &data); err != nil {
		return nil, statusCode, err
	}

	return &data, statusCode, nil
}

func (v *VerihubSdk) DetectFaceLiveness(face models.Face) (*models.FaceDetectResponse, *int, error) {
	mode := ""

	if v.Mode == 0 {
		mode = ""
	}

	parse, err := json.Marshal(&face)
	if err != nil {
		return nil, nil, err
	}

	body, statusCode, err := v.clientRequest(mode, string(parse), "/face/liveness", "v1")
	if err != nil {
		return nil, statusCode, err
	}

	var data models.FaceDetectResponse

	if err := json.Unmarshal(body, &data); err != nil {
		return nil, statusCode, err
	}

	return &data, statusCode, nil
}

func (v *VerihubSdk) SendSMSOTP(smsOtp models.RequestSmsOtp) (*models.ResponseSendOtp, *int, error) {
	mode := ""

	if v.Mode == 0 {
		mode = "/sandbox"
	}

	parse, err := json.Marshal(&smsOtp)
	if err != nil {
		return nil, nil, err
	}

	body, statusCode, err := v.clientRequest(mode, string(parse), "/otp/send", "v2")
	if err != nil {
		return nil, nil, err
	}

	var resOtp models.ResponseSendOtp

	if err := json.Unmarshal(body, &resOtp); err != nil {
		return nil, nil, err
	}

	return &resOtp, statusCode, nil
}

func (v *VerihubSdk) VerifyOTP(verify models.RequestVerifySmsOtp) (*models.ResponseMessage, *int, error) {
	mode := ""

	if v.Mode == 0 {
		mode = "/sandbox"
	}

	parse, err := json.Marshal(verify)
	if err != nil {
		return nil, nil, err
	}

	body, statusCode, err := v.clientRequest(mode, string(parse), "/otp/verify", "v2")
	if err != nil {
		return nil, nil, err
	}

	var resMessage models.ResponseMessage

	if err := json.Unmarshal([]byte(body), &resMessage); err != nil {
		return nil, nil, err
	}

	return &resMessage, statusCode, nil
}

func (v *VerihubSdk) ECIVerification(data models.VerificationData) (*interface{}, *int, error) {
	mode := ""

	if v.Mode == 0 {
		mode = "/sandbox"
	}

	parse, err := json.Marshal(data)
	if err != nil {
		return nil, nil, err
	}

	body, statusCode, err := v.clientRequest(mode, string(parse), "/data-verification/certificate-electronic", "v2")
	if err != nil {
		return nil, nil, err
	}

	var resData interface{}
	if err := json.Unmarshal([]byte(body), &resData); err != nil {
		return nil, nil, err
	}

	return &resData, statusCode, nil
}
