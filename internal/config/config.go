package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	DatabaseURI string `env:"DB_URI" env-required:"true"`
}

func New() (cfg *Config, err error) {
	return cfg, cleanenv.ReadConfig(".env", &cfg)
}
