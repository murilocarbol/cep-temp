package mocks

import (
	"github.com/murilocarbol/cep-temp/application/model"
	"github.com/stretchr/testify/mock"
)

type MockTemperatureUseCase struct {
	mock.Mock
}
type MockViaCepClient struct {
	mock.Mock
}
type MockWeatherClient struct {
	mock.Mock
}

func NewMockTemperatureUseCase() *MockTemperatureUseCase {
	return &MockTemperatureUseCase{}
}

func NewMockViaCepClient() *MockViaCepClient {
	return &MockViaCepClient{}
}

func NewMockWeatherClient() *MockWeatherClient {
	return &MockWeatherClient{}
}

func (c *MockTemperatureUseCase) GetTemperature(cep string) (*model.Temperature, error) {
	args := c.Called(cep)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Temperature), args.Error(1)
}

func (c *MockViaCepClient) GetEndereco(cep string) (string, error) {

	args := c.Called(cep)
	if args.Get(0) == nil {
		return "", args.Error(1)
	}
	return args.Get(0).(string), args.Error(1)
}

func (c *MockWeatherClient) GetWeather(localitation string) (*model.Temperature, error) {

	args := c.Called(localitation)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Temperature), args.Error(1)
}
