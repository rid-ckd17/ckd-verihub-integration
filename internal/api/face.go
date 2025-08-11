package api

import (
	"context"
	"errors"
	"testapiverihub/internal/models"
	"testapiverihub/internal/services"

	"github.com/danielgtaylor/huma/v2"
	"github.com/gofiber/fiber/v2"
)

type FaceHandler struct {
	FaceService *services.FaceService
	Ctx         *context.Context
}

func NewFaceHandler(faceSer *services.FaceService, ctx *context.Context) *FaceHandler {
	return &FaceHandler{FaceService: faceSer, Ctx: ctx}
}

func (h *FaceHandler) Route(api fiber.Router, appx huma.API) {
	api.Post("/faceenroll", h.EnrollFace)
	api.Post("/faceliveness", h.DetectFaceLiveness)

	huma.Post(appx, "/faceenroll", h.HumaEnrollFace)
	huma.Post(appx, "/faceliveness", h.HumaDetectFaceLiveness)
}

func (h *FaceHandler) EnrollFace(c *fiber.Ctx) error {
	var face models.Face
	if err := c.BodyParser(&face); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	message, statusCode, err := h.FaceService.EnrollFace(face)
	if err != nil {
		return c.Status(*statusCode).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(message)
}

func (h *FaceHandler) DetectFaceLiveness(c *fiber.Ctx) error {
	var face models.Face
	if err := c.BodyParser(&face); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	message, statusCode, err := h.FaceService.DetectFaceLiveness(face)
	if err != nil {
		return c.Status(*statusCode).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(message)
}

func (h *FaceHandler) HumaEnrollFace(c context.Context, input *models.Face) (*models.FaceDataResponse, error) {

	message, statusCode, err := h.FaceService.EnrollFace(*input)
	if err != nil {
		return nil, err
	}

	code := *statusCode

	if code == 400 || code == 403 || code == 409 || code == 422 || code == 429 || code == 500 {
		return nil, huma.NewError(code, "Kesalahan dari service", errors.New("wrong way"))
	}

	return message, nil
}

func (h *FaceHandler) HumaDetectFaceLiveness(c context.Context, input *models.Face) (*models.FaceDetectResponse, error) {

	message, statusCode, err := h.FaceService.DetectFaceLiveness(*input)
	if err != nil {
		return nil, err
	}

	code := *statusCode

	if code == 400 || code == 403 || code == 409 || code == 422 || code == 429 || code == 500 {
		return nil, huma.NewError(code, "Kesalahan dari service", errors.New("wrong way"))
	}

	return message, nil
}
