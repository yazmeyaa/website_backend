package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	JWT    JWTConfig    `yaml:"jwt"`
	Server ServerConfig `yaml:"server"`
	Redis  RedisConfig  `yaml:"redis"`
}

type JWTConfig struct {
	Secret string `yaml:"secret_key"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}

type RedisConfig struct {
	Database int    `yaml:"database"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func readConfig() *AppConfig {
	file, err := os.Open("config.yml")
	if err != nil {
		processError(err)
	}
	defer file.Close()

	var cfg AppConfig
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		processError(err)
	}

	return &cfg
}

func NewAppConfig() *AppConfig {
	config := readConfig()
	fmt.Printf("%+v", config)
	return config
}
