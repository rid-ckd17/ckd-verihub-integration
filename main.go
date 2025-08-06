package main

import (
	"context"
	"testapiverihub/internal/api"
	verihub "testapiverihub/internal/clients/verihubs"
	"testapiverihub/internal/services"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

const (
	Sandbox    = 0
	Production = 1
)

const (
	AppID  = "7ec45ffe-9a36-4756-8489-cc560b42b6c3"
	ApiKey = "3a5BYdgtJOYNDtKRx3NqnXbvGWGSH/qz"
	UrlV1  = "https://api.verihubs.com/v1"
	UrlV2  = "https://api.verihubs.com/v2"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	app := fiber.New()
	app.Use(logger.New())

	apiG := app.Group("/api")
	verihubClient := verihub.NewVirehubSdk(AppID, ApiKey, &ctx, Sandbox, UrlV1, UrlV2)

	smsService := services.NewSmsOtpService(verihubClient, &ctx)
	smsHandler := api.NewSmsOTPHandler(smsService, &ctx)

	eciService := services.NewECIService(verihubClient, &ctx)
	eciHandler := api.NewECIHandler(eciService, &ctx)

	smsHandler.Route(apiG)
	eciHandler.Route(apiG)

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
