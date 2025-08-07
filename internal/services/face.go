package services

import (
	"context"
	verihub "testapiverihub/internal/clients/verihubs"
	"testapiverihub/internal/models"
)

type FaceService struct {
	VeriHub *verihub.VerihubSdk
	Ctx     *context.Context
}

func NewFaceServer(verihub *verihub.VerihubSdk, ctx *context.Context) *FaceService {
	return &FaceService{VeriHub: verihub, Ctx: ctx}
}

func (s *FaceService) EnrollFace(face models.Face) (*models.FaceDataResponse, *int, error) {
	data, statusCode, err := s.VeriHub.EnrollFace(face)
	if err != nil {
		return nil, statusCode, err
	}

	return data, statusCode, nil
}

func (s *FaceService) DetectFaceLiveness(face models.Face) (*models.FaceDetectResponse, *int, error) {
	data, statusCode, err := s.VeriHub.DetectFaceLiveness(face)
	if err != nil {
		return nil, statusCode, err
	}

	return data, statusCode, nil
}
