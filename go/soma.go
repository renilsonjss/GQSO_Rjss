package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func somaHandler(c *fiber.Ctx) error{
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
    total := soma(valor1, valor2)
    return c.SendString(fmt.Sprintf("%.2f", total))
}

func soma(op1str, op2str string) (float64, error) {
    op1, err := strconv.ParseFloat(op1str, 64)
	if err != nil {
		return 0, fmt.Errorf("Parâmetro inválido (%s):%q", op1, err)
	}
	op2, err := strconv.ParseFloat(op2str, 64)
	if err != nil {
		return 0, fmt.Errorf("Parâmetro inválido (%s):%q", op2, err)
	}
    return (op1 + op2), nil
}