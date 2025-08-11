package api

import (
	"context"
	"testapiverihub/internal/models"
	"testapiverihub/internal/services"

	"github.com/gofiber/fiber/v2"
)

type KTPExtractHandler struct {
	KTPExtractService *services.KTPExtractService
	Ctx               *context.Context
}

func NewKTPExtractHandler(ktpExtractService *services.KTPExtractService, ctx *context.Context) *KTPExtractHandler {
	return &KTPExtractHandler{KTPExtractService: ktpExtractService, Ctx: ctx}
}

func (h *KTPExtractHandler) Route(api fiber.Router) {
	api.Post("/ktpextract", h.ExtractKTP)
}

func (h *KTPExtractHandler) ExtractKTP(c *fiber.Ctx) error {
	var ktpExtract models.KTPExtract
	if err := c.BodyParser(&ktpExtract); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	data, statusCode, err := h.KTPExtractService.ExtractKTP(ktpExtract)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to extract KTP",
		})
	}

	code := *statusCode

	if code == 400 || code == 403 || code == 409 || code == 422 || code == 429 || code == 500 {
		return c.Status(code).JSON(data)
	}

	return c.JSON(data)
}
