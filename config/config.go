package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	TelegramBotToken string  `yaml:"telegram_token"`
	APIBaseURL       string  `yaml:"api_base_url"`
	AdminUsers       []int64 `yaml:"admin_users"`
	APILogin         string  `yaml:"login"`
	APIPassword      string  `yaml:"password"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	log.Println("config loaded successfully")
	return &cfg, nil
}
