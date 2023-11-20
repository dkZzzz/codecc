package config

import (
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	OpenAISecretKey string `json:"sercret_key"`
}

func init() {
}

func LoadConfig() (*Config, error) {
	config := &Config{}

	file, err := os.Open("config.json")

	if err != nil {
		return config, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
