package grpc_server

import "time"

type Config struct {
	Port    string        `envconfig:"GRPC_PORT"    required:"true"`
	Timeout time.Duration `envconfig:"GRPC_TIMEOUT" required:"true"`
}
