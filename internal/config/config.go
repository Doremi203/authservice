package config

import (
	"flag"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env  string     `yaml:"env" env-default:"local"`
	DB   DBConfig   `yaml:"db" env-required:"true"`
	Grpc GrpcConfig `yaml:"grpc" env-required:"true"`
	Auth AuthConfig `yaml:"auth" env-required:"true"`
}

type DBConfig struct {
	UserID         string `yaml:"user-id" env-required:"true"`
	Password       string `yaml:"password" env-required:"true"`
	Database       string `yaml:"database" env-required:"true"`
	Host           string `yaml:"host" env-required:"true"`
	Options        string `yaml:"options" env-required:"true"`
	MigrationsPath string `yaml:"migrations-path" env-required:"true"`
}

func (c DBConfig) ConnectionString() string {
	return fmt.Sprintf("User ID=%s;Password=%s;Host=%s;Database=%s;%s", c.UserID, c.Password, c.Host, c.Database, c.Options)
}

type GrpcConfig struct {
	Port    int           `yaml:"port" env-required:"true"`
	Timeout time.Duration `yaml:"timeout" env-required:"true"`
}

type AuthConfig struct {
	TokenTTL time.Duration `yaml:"token-ttl" env-required:"true"`
}

func MustLoad() Config {
	path := getConfigPath()
	if path == "" {
		panic("config file path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file not found at: " + path)
	}

	var cfg Config
	err := cleanenv.ReadConfig(path, &cfg)

	if err != nil {
		panic("failed to read config: " + err.Error())
	}

	return cfg
}

func getConfigPath() string {
	var path string

	flag.StringVar(&path, "config", "", "path to config file")
	flag.Parse()

	if path == "" {
		path = os.Getenv("CONFIG_PATH")
	}

	return path
}
