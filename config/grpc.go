package config

import "os"

const (
	defaultGRPCServHost = "127.0.0.1"
	defaultGRPCServPort = "6060"
)
const (
	envNameGRPCServHost = "GRPC_HOST"
	envNameGRPCServPort = "GRPC_PORT"
)

type GRPCConfig struct {
	Host string
	Port string
}

func newDefaultGRPCConfig() GRPCConfig {
	return GRPCConfig{
		Host: defaultGRPCServHost,
		Port: defaultGRPCServPort,
	}
}

func (c *GRPCConfig) parseEnv() {
	envServHost := os.Getenv(envNameGRPCServHost)
	if envServHost != "" {
		c.Host = envServHost
	}

	envServPort := os.Getenv(envNameGRPCServPort)
	if envServPort != "" {
		c.Port = envServPort
	}
}
