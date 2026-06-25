package config

type HTTP struct {
	Host    string `env:"NODE_HOST" env-default:"localhost"`
	Port    int    `env:"NODE_PORT" env-default:"3000"`
	Timeout int    `env:"NODE_STARTING_TIMEOUT_MS" env-default:"100"`
}
