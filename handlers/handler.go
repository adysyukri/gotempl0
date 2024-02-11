package handlers

import "github.com/gofiber/fiber/v2"

func NewHandler(staticPath string) *fiber.App {
	app := fiber.New()

	setupRoutes(app, staticPath)

	return app
}
