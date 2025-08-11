package services

import (
	"context"
	"testapiverihub/internal/models"
)

type VerihubSdkInterface interface {
	EnrollFace(face models.Face) (*models.FaceDataResponse, *int, error)
	DetectFaceLiveness(face models.Face) (*models.FaceDetectResponse, *int, error)
}

type FaceService struct {
	VeriHub VerihubSdkInterface
	Ctx     *context.Context
}

func NewFaceServer(verihub VerihubSdkInterface, ctx *context.Context) *FaceService {
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
