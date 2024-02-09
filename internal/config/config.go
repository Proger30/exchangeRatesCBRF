package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	UpdateTime int          `json:"update_time"`
	Currencies []string     `json:"currencies"`
	DB         DBConfig     `json:"db"`
	Server     ServerConfig `json:"server"`
}

type DBConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
}

type ServerConfig struct {
	Port int `json:"port"`
}

func LoadConfig(filename string) (Config, error) {
	var cfg Config

	data, err := os.ReadFile(filename)
	if err != nil {
		return cfg, err
	}

	if err := json.Unmarshal(data, &cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}
