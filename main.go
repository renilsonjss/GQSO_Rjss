package main

import (
	"github.com/gofiber/fiber/v2"
)

//Para utilização da funcionalidade soma no navegador:
    //1 - Impementar o comando "go build && ./calc-server" no terminal (sem as aspas);
    //2 - Selecionar a opção abrir no navegador (open in browser);
    //3 - Inserir, no fim da URL, "soma/'valor1'/'valor2'" (sem as aspas);
    //4 - O resultado da soma de "valor 1" e "valor 2" será exibid.

func main(){
	app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Para utilização da funcionalidade soma, inserir, no fim da URL: 'soma/valor1/valor2' (sem as aspas).")
    })

	app.Get("/soma/:valor1/:valor2", soma)

    app.Listen(":8000")
}