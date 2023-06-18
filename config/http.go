package config

import "os"

const (
	defaultHTTPServHost = "127.0.0.1"
	defaultHTTPServPort = "6969"
)
const (
	envNameHTTPServHost = "HTTP_HOST"
	envNameHTTPServPort = "HTTP_PORT"
)

type HTTPConfig struct {
	Host string
	Port string
}

func newDefaultHTTPConfig() HTTPConfig {
	return HTTPConfig{
		Host: defaultHTTPServHost,
		Port: defaultHTTPServPort,
	}
}

func (c *HTTPConfig) parseEnv() {
	envServHost := os.Getenv(envNameHTTPServHost)
	if envServHost != "" {
		c.Host = envServHost
	}

	envServPort := os.Getenv(envNameHTTPServPort)
	if envServPort != "" {
		c.Port = envServPort
	}
}
