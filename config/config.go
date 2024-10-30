package config

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/murilocarbol/cep-temp/application/client"
	"github.com/murilocarbol/cep-temp/application/controllers"
	"github.com/murilocarbol/cep-temp/application/usecases"
	"github.com/spf13/viper"
)

type configure struct {
	WEATHER_API_KEY string `mapstructure:"WEATHER_API_KEY"`
}

func Initialize() {

	confg, _ := LoadConfig(".")

	app := fiber.New()
	setRoutes(app, confg.WEATHER_API_KEY)
	app.Listen(":8080")
}

func setRoutes(app *fiber.App, key string) {
	// Clients
	viaCepClient := client.NewViaCepClient()
	weatherClient := client.NewWeatherClient(key)

	// Usecases
	temperatureUseCase := usecases.NewTemperatureUseCase(viaCepClient, weatherClient)

	// Controllers
	temperatureController := controllers.NewTemperatureController(temperatureUseCase)

	app.Get("/", temperatureController.GetTemperature)
}

func LoadConfig(path string) (*configure, error) {
	var cfg *configure
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath("./")
	viper.SetConfigFile("config.env")
	viper.AutomaticEnv()

	fmt.Println("Loading config from path:", path)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg, err
}
