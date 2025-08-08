package api

import (
	"context"
	"testapiverihub/internal/models"
	"testapiverihub/internal/services"

	"github.com/danielgtaylor/huma/v2"
	"github.com/gofiber/fiber/v2"
)

type SmsOTPHandler struct {
	SmSOTPService *services.SmsOTPService
	Ctx           *context.Context
}

func NewSmsOTPHandler(smsOtpService *services.SmsOTPService, ctx *context.Context) *SmsOTPHandler {
	return &SmsOTPHandler{SmSOTPService: smsOtpService, Ctx: ctx}
}

func (h *SmsOTPHandler) Route(api fiber.Router, ap huma.API) {
	api.Post("/sendotp", h.SendOTP)
	api.Post("/verifyotp", h.VeriFyOTP)
}

func (h *SmsOTPHandler) SendOTP(c *fiber.Ctx) error {
	var sendData models.RequestSmsOtp

	if err := c.BodyParser(&sendData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	message, statusCode, err := h.SmSOTPService.SendSMSOTP(sendData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	code := *statusCode

	if code == 400 || code == 403 || code == 409 || code == 422 || code == 429 || code == 500 {
		return c.Status(code).JSON(message)
	}

	return c.JSON(message)
}

func (h *SmsOTPHandler) VeriFyOTP(c *fiber.Ctx) error {
	var verifyData models.RequestVerifySmsOtp

	if err := c.BodyParser(&verifyData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	message, statusCode, err := h.SmSOTPService.VerifySMSOTP(verifyData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	code := *statusCode

	if code == 400 || code == 403 || code == 409 || code == 422 || code == 429 || code == 500 {
		return c.Status(code).JSON(message)
	}

	return c.JSON(message)
}
