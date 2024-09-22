package cfg

import (
	"errors"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	GRPC    GRPCConfig
	MongoDB MongoDBConfig
}

type GRPCConfig struct {
	Port string
}

type MongoDBConfig struct {
	Port       string
	Host       string
	DBName     string
	Collection string
}

func parseConfig(v *viper.Viper) (*Config, error) {
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		log.Printf("unable to parse config: %v", err)
		return nil, err
	}
	return &cfg, nil
}

func GetConfig() *Config {
	cfgPath := getConfigPath(os.Getenv("APP_ENV"))

	v, err := LoadConfig(cfgPath, "yaml")

	if err != nil {
		log.Fatalf("unable to get config: %v", err)
	}

	cfg, err := parseConfig(v)

	if err != nil {
		log.Fatalf("unable to parse config: %v", err)
	}

	return cfg
}

func LoadConfig(path string, fileType string) (*viper.Viper, error) {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigType(fileType)
	v.SetConfigName(path)
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		log.Printf("unable to load config: %v", err)
		var viperErr viper.ConfigFileNotFoundError
		if errors.As(err, &viperErr) {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

func getConfigPath(env string) string {
	if env == "local" {
		return "/config/config-local.yaml"
	} else if env == "production" {
		return "/config/config-prod.yaml"
	} else {
		return "/config/config-dev.yaml"
	}
}
