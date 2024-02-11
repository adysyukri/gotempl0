package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/adysyukri/gotempl0/templates"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func setupRoutes(app *fiber.App, staticPath string) {
	app.Get("/", func(c *fiber.Ctx) error {
		return render(c, templates.Index("Home root", templates.Hello("Ajim")))
	})
	app.Static("/static", "./static")

	app.Use(notFoundMiddleware)

}

func notFoundMiddleware(c *fiber.Ctx) error {
	return render(c, templates.Index("Not Found", templates.NotFound()), templ.WithStatus(http.StatusNotFound))
}

func render(c *fiber.Ctx, component templ.Component, options ...func(*templ.ComponentHandler)) error {
	componentHandler := templ.Handler(component)
	for _, o := range options {
		o(componentHandler)
	}
	return adaptor.HTTPHandler(componentHandler)(c)
}
