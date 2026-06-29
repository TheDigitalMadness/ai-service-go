package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	HTTP HTTP
	AI   AI
}

func MustLoad() *Config {
	cfg := Config{}

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		panic(err)
	}

	fmt.Println("Config loaded successfully")

	return &cfg
}
