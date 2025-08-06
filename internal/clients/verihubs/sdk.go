package verihub

import (
	"context"
	"encoding/json"
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

func (v *VerihubSdk) ClientRequest(data string, path string) ([]byte, *int, error) {
	mode := ""

	if v.Mode == 0 {
		mode = "/sandbox"
	}

	url := v.URLV2 + path + mode
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
	return body, &res.StatusCode, nil
}

func (v *VerihubSdk) SendSMSOTP(smsOtp models.RequestSmsOtp) (*models.ResponseSendOtp, *int, error) {
	parse, err := json.Marshal(&smsOtp)
	if err != nil {
		return nil, nil, err
	}

	body, statusCode, err := v.ClientRequest(string(parse), "/otp/send")
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
	parse, err := json.Marshal(verify)
	if err != nil {
		return nil, nil, err
	}

	body, statusCode, err := v.ClientRequest(string(parse), "/otp/verify")
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
	parse, err := json.Marshal(data)
	if err != nil {
		return nil, nil, err
	}

	body, statusCode, err := v.ClientRequest(string(parse), "/data-verification/certificate-electronic")
	if err != nil {
		return nil, nil, err
	}

	var resData interface{}
	if err := json.Unmarshal([]byte(body), &resData); err != nil {
		return nil, nil, err
	}

	return &resData, statusCode, nil
}
