package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Storage string `yaml:"storage" env-required:"true"`
	Log     string `yaml:"log_level" env-default:"debug"`
	URL     string `env:"URL"`
}

func New() *Config {
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loadnig .env file. %v", err)
	}
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		logrus.Fatalln("config path is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		logrus.Fatalf("config file does not exists. %v", configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		logrus.Fatalf("cannot read config. %v", err)
	}
	return &cfg

}
