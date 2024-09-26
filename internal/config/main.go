package config

import "github.com/updevru/go-micro-kit/config"

type Config struct {
	config.App
	Http    config.Http `env:",prefix=HTTP_"`
	Grpc    config.Grpc `env:",prefix=GRPC_"`
	Cluster Cluster     `env:",prefix=CLUSTER_"`
}

type Cluster struct {
	Servers []string `env:"SERVERS" envSeparator:","`
}
