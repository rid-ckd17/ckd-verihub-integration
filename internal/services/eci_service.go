package services

import (
	"context"
	verihub "testapiverihub/internal/clients/verihubs"
	"testapiverihub/internal/models"
)

type ECIService struct {
	VeriHub *verihub.VerihubSdk
	Ctx     *context.Context
}

func NewECIService(verihub *verihub.VerihubSdk, ctx *context.Context) *ECIService {
	return &ECIService{VeriHub: verihub, Ctx: ctx}
}

func (s *ECIService) Verification(verData models.VerificationData) (*interface{}, *int, error) {
	resData, statusCode, err := s.VeriHub.ECIVerification(verData)
	if err != nil {
		return nil, nil, err
	}
	return resData, statusCode, nil
}
