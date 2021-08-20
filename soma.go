package main

import "fmt"

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func soma(c *fiber.Ctx) error{
    valor1Str:=c.Params("valor1")
    valor1, err:= strconv.ParseFloat(valor1Str, 64)
    if err != nil{
        return c.Status(http.StatusBadRequest).SendString(fmt.Sprintf("Parâmetro inválido:\":%s\"", valor1Str))
    }

    valor2Str:=c.Params("valor2")
    valor2, err:= strconv.ParseFloat(valor2Str, 64)
    if err != nil{
        return c.Status(http.StatusBadRequest).SendString(fmt.Sprintf("Parâmetro inválido:\":%s\"", valor2Str))
    }
    total := valor1 + valor2
    return c.SendString(fmt.Sprintf("%.2f", total))
}