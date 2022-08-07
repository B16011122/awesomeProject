package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Match any route
	app.Use(func(c *fiber.Ctx) error {
		fmt.Println("First middleware")
		c.Next()
		return nil
	})

	// Match all routes starting with /api
	app.Use("/api", func(c *fiber.Ctx) error {
		fmt.Println("Second middleware")
		c.Next()
		return nil
	})

	// GET /api/register
	app.Get("/api/list", func(c *fiber.Ctx) error {
		fmt.Println("Last middleware")
		c.Send([]byte("Hello, World!"))
		return nil
	})
	app.Listen(":3000")
}
