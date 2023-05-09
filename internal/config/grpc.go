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
	ServHost string
	ServPort string
}

func newDefaultGRPCConfig() GRPCConfig {
	return GRPCConfig{
		ServHost: defaultGRPCServHost,
		ServPort: defaultGRPCServPort,
	}
}

func (c *GRPCConfig) parseEnv() {
	envServHost := os.Getenv(envNameGRPCServHost)
	if envServHost != "" {
		c.ServHost = envServHost
	}

	envServPort := os.Getenv(envNameGRPCServPort)
	if envServPort != "" {
		c.ServPort = envServPort
	}
}
