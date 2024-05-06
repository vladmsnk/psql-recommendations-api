package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"psqlRecommendationsApi/lib/grpc_server"
)

const pathToConfig = "config/config.yaml"

type Config struct {
	GRPC              grpc_server.GRPCConfig `yaml:"grpc"`
	PG                Postgres               `yaml:"postgres"`
	Collector         Collector              `yaml:"collector"`
	RecommendationApi RecommendationApi      `yaml:"recommendation_api"`
}

type Postgres struct {
	Host          string `yaml:"host"`
	Port          int    `yaml:"port"`
	User          string `yaml:"user"`
	Password      string `yaml:"password"`
	Database      string `yaml:"database"`
	SSLMode       string `yaml:"sslmode"`
	ContainerName string `yaml:"container_name"`
}

type Collector struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type RecommendationApi struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

var ConfigStruct Config

func Init() error {
	rawYaml, err := os.ReadFile(pathToConfig)
	if err != nil {
		return fmt.Errorf("os.ReadFile: %w", err)
	}

	if err = yaml.Unmarshal(rawYaml, &ConfigStruct); err != nil {
		return fmt.Errorf("yaml.Unmarshal: %w", err)
	}
	return nil
}
