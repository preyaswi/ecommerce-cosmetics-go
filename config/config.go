package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	BASE_URL   string `mapstructure:"BASE_URL"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBName     string `mapstructure:"DB_NAME"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBPassword string `mapstructure:"DB_PASSWORD"`

	AUTHTOKEN   string `mapstructure:"TWILIO_AUTHTOKEN"`
	ACCOUNTSID  string `mapstructure:"TWILIO_ACCOUNTSID"`
	SERVICESSID string `mapstructure:"TWILIO_SERVICESID"`

	KEY                string `mapstructure:"KEY"`
	KEY_FOR_ADMIN      string `mapstructure:"KEY_FOR_ADMIN"`
	KEY_ID_fOR_PAY     string `mapstructure:"KEY_ID_fOR_PAY"`
	SECRET_KEY_FOR_PAY string `mapstructure:"SECRET_KEY_FOR_PAY"`

	API_KEY_FOR_CLOUDINARY    string `mapstructure:"API_KEY_FOR_CLOUDINARY"`
	API_SECRET_FOR_CLOUDINARY string `mapstructure:"API_SECRET_FOR_CLOUDINARY"`
	CLOUD_NAME                string `mapstructure:"CLOUD_NAME"`
}

var envs = []string{
	"BASE_URL", "DB_HOST", "DB_NAME", "DB_USER", "DB_PORT", "DB_PASSWORD", "TWILIO_AUTHTOKEN", "TWILIO_ACCOUNTSID", "TWILIO_SERVICESID", "KEY_FOR_ADMIN", "KEY", "KEY_ID_fOR_PAY", "SECRET_KEY_FOR_PAY", "API_KEY_FOR_CLOUDINARY", "API_SECRET_FOR_CLOUDINARY", "CLOUD_NAME",
}

func LoadConfig() (Config, error) {
	var config Config

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}
	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}
	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}
	return config, nil

}
