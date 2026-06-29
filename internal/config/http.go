package config

type HTTP struct {
	Port    int `env:"NODE_PORT" env-default:"3000"`
	Timeout int `env:"NODE_STARTING_TIMEOUT_MS" env-default:"100"`
}
