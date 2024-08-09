package token

import "time"

type Config struct {
	TokenTTL time.Duration `yaml:"ttl" env-required:"true"`
	Key      string        `yaml:"key" env:"KEY" env-required:"true"`
}
