package config

import "github.com/updevru/go-micro-kit/config"

type Config struct {
	config.App
	Http    config.Http `env:",prefix=HTTP_"`
	Grpc    config.Grpc `env:",prefix=GRPC_"`
	Cluster Cluster     `env:",prefix=CLUSTER_"`
	Storage Storage     `env:",prefix=STORAGE_"`
}

type Cluster struct {
	Servers []string `env:"SERVERS" envSeparator:","`
}

type Storage struct {
	Name string      `env:"NAME, required"`
	Bolt StorageBolt `env:",prefix=BOLT_"`
}

type StorageBolt struct {
	File string `env:"FILE"`
}
