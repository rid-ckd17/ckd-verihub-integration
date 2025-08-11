package services

import (
	"context"
	"testing"

	"testapiverihub/internal/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockVerihubSdk struct {
	mock.Mock
}

func (m *MockVerihubSdk) EnrollFace(face models.Face) (*models.FaceDataResponse, *int, error) {
	args := m.Called(face)
	return args.Get(0).(*models.FaceDataResponse), args.Get(1).(*int), args.Error(2)
}

func (m *MockVerihubSdk) DetectFaceLiveness(face models.Face) (*models.FaceDetectResponse, *int, error) {
	args := m.Called(face)
	return args.Get(0).(*models.FaceDetectResponse), args.Get(1).(*int), args.Error(2)
}

func TestEnrollFace(t *testing.T) {
	mockSdk := new(MockVerihubSdk)
	ctx := context.Background()
	service := NewFaceServer(mockSdk, &ctx) // Ensure NewFaceServer accepts an interface implemented by MockVerihubSdk

	face := models.Face{}
	response := &models.FaceDataResponse{}
	statusCode := new(int)
	*statusCode = 200

	mockSdk.On("EnrollFace", face).Return(response, statusCode, nil)

	result, status, err := service.EnrollFace(face)

	assert.NoError(t, err)
	assert.Equal(t, response, result)
	assert.Equal(t, statusCode, status)
	mockSdk.AssertExpectations(t)
}

func TestDetectFaceLiveness(t *testing.T) {
	mockSdk := new(MockVerihubSdk)
	ctx := context.Background()
	service := NewFaceServer(mockSdk, &ctx) // NewFaceServer should accept an interface, not a concrete type

	face := models.Face{}
	response := &models.FaceDetectResponse{}
	statusCode := new(int)
	*statusCode = 200

	mockSdk.On("DetectFaceLiveness", face).Return(response, statusCode, nil)

	result, status, err := service.DetectFaceLiveness(face)

	assert.NoError(t, err)
	assert.Equal(t, response, result)
	assert.Equal(t, statusCode, status)
	mockSdk.AssertExpectations(t)
}
