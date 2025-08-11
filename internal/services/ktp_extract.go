package services

import (
	"context"
	verihub "testapiverihub/internal/clients/verihubs"
	"testapiverihub/internal/models"
)

type KTPExtractService struct {
	VeriHub *verihub.VerihubSdk
	Ctx     *context.Context
}

func NewKTPExtractService(verihub *verihub.VerihubSdk, ctx *context.Context) *KTPExtractService {
	return &KTPExtractService{VeriHub: verihub, Ctx: ctx}
}

func (s *KTPExtractService) ExtractKTP(ktpExtract models.KTPExtract) (*interface{}, *int, error) {
	data, statusCode, err := s.VeriHub.ExtractAsyncKTP(ktpExtract)
	if err != nil {
		return nil, statusCode, err
	}
	return data, statusCode, nil
}
