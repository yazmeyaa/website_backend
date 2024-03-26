package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	JWT struct {
		Secret string `yaml:"secret_key"`
	} `yaml:"jwt"`
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
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
