package postgres

import "fmt"

type Config struct {
	UserID         string `yaml:"user-id" env-required:"true"`
	Password       string `yaml:"password" env-required:"true"`
	Database       string `yaml:"database" env-required:"true"`
	Host           string `yaml:"host" env-required:"true"`
	Options        string `yaml:"options" env-required:"true"`
	MigrationsPath string `yaml:"migrations-path" env-required:"true"`
}

func (c Config) ConnectionString() string {
	return fmt.Sprintf("User ID=%s;Password=%s;Host=%s;Database=%s;%s", c.UserID, c.Password, c.Host, c.Database, c.Options)
}
