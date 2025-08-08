package api

import (
	"context"
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
