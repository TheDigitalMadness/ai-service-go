package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	HTTP HTTP
	AI   AI
}

func MustLoad() *Config {
	cfg := Config{}

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		panic(err)
	}

	return &cfg
}
