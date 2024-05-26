package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func main() {
	app := fiber.New()

	app.Get("/:name?", func(c *fiber.Ctx) error {
		name := c.Params("name")
		if name == "" {
			name = "World"
		}
		return Render(c, Home(name))
	})
	app.Use(NotFoundMiddleware)

	log.Fatal(app.Listen(":3000"))
}

func NotFoundMiddleware(c *fiber.Ctx) error {
	return Render(c, NotFound(), t1.WithStatus(http.StatusNotFound))
}

func Render(c *fiber.Ctx, component t1.Component, options ...func(*t1.ComponentHandler)) error {
	componentHandler := t1.Handler(component)
	for _, o := range options {
		o(componentHandler)
	}
	return adaptor.HTTPHandler(componentHandler)(c)
}
