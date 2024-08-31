package postgres

import "fmt"

type Config struct {
	UserID         string `yaml:"user-id" env-required:"true"`
	Password       string `yaml:"password" env-required:"true"`
	Database       string `yaml:"database" env-required:"true"`
	Host           string `yaml:"host" env-required:"true"`
	Port           int    `yaml:"port" env-required:"true"`
	MigrationsPath string `yaml:"migrations-path" env-required:"true"`
}

func (c Config) ConnectionString() string {
	return fmt.Sprintf("user=%s password=%s host=%s database=%s port=%d sslmode=disable", c.UserID, c.Password, c.Host, c.Database, c.Port)
}
