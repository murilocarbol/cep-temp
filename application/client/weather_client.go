package client

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/murilocarbol/cep-temp/application/client/response"
	"github.com/murilocarbol/cep-temp/application/model"
)

type WeatherClient struct {
	api_key string
}

func NewWeatherClient(key string) *WeatherClient {
	return &WeatherClient{
		api_key: key,
	}
}

type WeatherClientInterface interface {
	GetWeather(localitation string) (*model.Temperature, error)
}

func (v WeatherClient) GetWeather(localitation string) (*model.Temperature, error) {

	req, err := http.NewRequest("GET", "http://api.weatherapi.com/v1/current.json", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("key", v.api_key)
	q.Add("q", localitation)
	q.Add("aqi", "no")
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var weather response.Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		return nil, err
	}

	temperature := &model.Temperature{
		TempC: weather.Current.TempC,
		TempF: weather.Current.TempF,
	}

	return temperature, nil
}
