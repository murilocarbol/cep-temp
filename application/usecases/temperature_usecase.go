package usecases

import (
	"fmt"
	"log"

	"github.com/murilocarbol/cep-temp/application/client"
	"github.com/murilocarbol/cep-temp/application/model"
)

type TemperatureUseCase struct {
	viaCepClient  client.ViaCepClient
	weatherClient client.WeatherClient
}

func NewTemperatureUseCase(viaCepClient *client.ViaCepClient, weatherClient *client.WeatherClient) *TemperatureUseCase {
	return &TemperatureUseCase{
		viaCepClient:  *viaCepClient,
		weatherClient: *weatherClient,
	}
}

func (t *TemperatureUseCase) GetTemperature(cep string) (*model.Temperature, error) {

	city, err := t.viaCepClient.GetEndereco(cep)
	if err != nil {
		return nil, fmt.Errorf("zipcode not found")
	}

	log.Printf("City: %s", city)

	temp, err := t.weatherClient.GetWeather(city)
	if err != nil {
		return nil, err
	}

	temp.TempK = temp.TempC + 273

	return temp, nil
}
