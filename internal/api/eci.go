package api

import (
	"context"
	"testapiverihub/internal/models"
	"testapiverihub/internal/services"

	"github.com/gofiber/fiber/v2"
)

type ECIHandler struct {
	ECIService *services.ECIService
	Ctx        *context.Context
}

func NewECIHandler(eciser *services.ECIService, ctx *context.Context) *ECIHandler {
	return &ECIHandler{ECIService: eciser, Ctx: ctx}
}

func (h *ECIHandler) Route(api fiber.Router) {
	api.Post("/eci-verify", h.VerifyECI)
}

func (h *ECIHandler) VerifyECI(c *fiber.Ctx) error {
	var verData models.VerificationData
	if err := c.BodyParser(&verData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	message, statusCode, err := h.ECIService.Verification(verData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	code := *statusCode

	if code == 400 || code == 403 || code == 409 || code == 422 || code == 429 || code == 500 {
		return c.Status(fiber.StatusBadRequest).JSON(message)
	}

	return c.JSON(message)
}
