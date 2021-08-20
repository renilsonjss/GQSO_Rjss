package main

import (
	"github.com/gofiber/fiber/v2"
)

func main(){
	app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })

	app.Get("/soma/:valor1/:valor2", soma)

    app.Listen(":8000")
}