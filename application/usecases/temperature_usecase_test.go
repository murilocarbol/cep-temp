package usecases

import (
	"testing"

	"github.com/murilocarbol/cep-temp/application/client"
	"github.com/murilocarbol/cep-temp/application/mocks"
	"github.com/murilocarbol/cep-temp/application/model"
	"github.com/stretchr/testify/assert"
)

func TestUseCase(t *testing.T) {

	cep := "13201005"
	localidade := "Jundiaí"
	temperatura := &model.Temperature{
		TempC: 24,
		TempF: 75.2,
		TempK: 298.15,
	}

	cv := mocks.NewMockViaCepClient()
	cw := mocks.NewMockWeatherClient()
	useCase := mocks.NewMockTemperatureUseCase()

	t.Run("GetTemperature", func(t *testing.T) {
		cv.On("GetEndereco", cep).Return(localidade, nil)
		cw.On("GetWeather", localidade).Return(temperatura, nil)
		useCase.On("GetTemperature", cep).Return(temperatura, nil)

		temperature, err := TemperatureUseCaseInterface.GetTemperature(useCase, cep)
		if err != nil {
			panic(err)
		}

		assert.Nil(t, err)
		assert.Equal(t, temperature.TempC, temperatura.TempC)
		assert.Equal(t, temperature.TempF, temperatura.TempF)
		assert.Equal(t, temperature.TempK, temperatura.TempK)
	})

	t.Run("GetEndereco", func(t *testing.T) {
		cv.On("GetEndereco", cep).Return(localidade, nil)

		city, err := client.ViaCepClientInterface.GetEndereco(cv, cep)
		if err != nil {
			panic(err)
		}

		assert.Nil(t, err)
		assert.Equal(t, city, localidade)
	})

	t.Run("GetTemperature", func(t *testing.T) {

		cw.On("GetWeather", localidade).Return(temperatura, nil)

		temp, err := client.WeatherClientInterface.GetWeather(cw, localidade)
		if err != nil {
			panic(err)
		}

		assert.Nil(t, err)
		assert.Equal(t, temp.TempC, temperatura.TempC)
		assert.Equal(t, temp.TempF, temperatura.TempF)
	})

	t.Run("GetTemperature", func(t *testing.T) {

	})
}
