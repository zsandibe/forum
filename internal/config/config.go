package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Port    string
	Migrate string
	DB      struct {
		Dsn    string
		Driver string
	}
}

func NewConfig() (Config, error) {
	// Открываем JSON-файл с конфигурацией.
	configFile, err := os.Open("config.json")
	if err != nil {
		return Config{}, err
	}
	defer configFile.Close()
	// Декодируем JSON-файл в структуру Config.
	var config Config
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
