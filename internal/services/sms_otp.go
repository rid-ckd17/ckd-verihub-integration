package services

import (
	"context"
	verihub "testapiverihub/internal/clients/verihubs"
	"testapiverihub/internal/models"
)

type SmsOTPService struct {
	VeriHub *verihub.VerihubSdk
	Ctx     *context.Context
}

func NewSmsOtpService(verihub *verihub.VerihubSdk, ctx *context.Context) *SmsOTPService {
	return &SmsOTPService{VeriHub: verihub, Ctx: ctx}
}

func (s *SmsOTPService) SendSMSOTP(sendData models.RequestSmsOtp) (*models.ResponseSendOtp, *int, error) {
	resOtp, statusCode, err := s.VeriHub.SendSMSOTP(sendData)
	if err != nil {
		return nil, nil, err
	}

	return resOtp, statusCode, nil
}

func (s *SmsOTPService) VerifySMSOTP(verifyData models.RequestVerifySmsOtp) (*models.ResponseMessage, *int, error) {
	resMess, statusCode, err := s.VeriHub.VerifyOTP(verifyData)
	if err != nil {
		return nil, nil, err
	}

	return resMess, statusCode, nil
}
