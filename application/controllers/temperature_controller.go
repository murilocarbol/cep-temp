package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/murilocarbol/cep-temp/application/controllers/response"
	"github.com/murilocarbol/cep-temp/application/usecases"
)

type TemperatureController struct {
	temperatureUseCase usecases.TemperatureUseCase
}

func NewTemperatureController(temperatureUsecase *usecases.TemperatureUseCase) *TemperatureController {
	return &TemperatureController{
		temperatureUseCase: *temperatureUsecase,
	}
}

func (c *TemperatureController) GetTemperature(ctx *fiber.Ctx) error {
	cep := ctx.Query("cep")

	if cep == "" || len(cep) != 8 {
		return ctx.Status(422).JSON(response.ErrorResponse{
			Error: "invalid zipcode",
		})
	}

	temperatures, err := c.temperatureUseCase.GetTemperature(cep)
	if err != nil {
		if err.Error() == "zipcode not found" {
			return ctx.Status(404).JSON(response.ErrorResponse{Error: err.Error()})
		}
		return ctx.Status(500).JSON(response.ErrorResponse{
			Error: "internal server error",
		})
	}

	log.Printf("Temperatures: %+v", temperatures)

	temperaturesResponse := &response.TemperatureResponse{
		TempC: temperatures.TempC,
		TempF: temperatures.TempF,
		TempK: temperatures.TempK,
	}

	return ctx.Status(200).JSON(temperaturesResponse)
}
