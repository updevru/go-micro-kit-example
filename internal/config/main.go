package config

import "github.com/updevru/go-micro-kit/config"

type Config struct {
	config.App
	Http config.Http `env:",prefix=HTTP_"`
	Grpc config.Grpc `env:",prefix=GRPC_"`
}
