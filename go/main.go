package main

import (
	"github.com/gofiber/fiber/v2"
)

func main(){
	app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Para utilização da funcionalidade soma, inserir, no fim da URL: 'soma/valor1/valor2' (sem as aspas).")
    })

	app.Get("/soma/:valor1/:valor2", somaHandler)

    app.Listen(":8000")
}