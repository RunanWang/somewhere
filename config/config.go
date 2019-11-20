package config

import (
	"github.com/BurntSushi/toml"
)

type TConfig struct {
	ServiceConfig TServiceConfig `toml:"server"`
	DbConfig      TDbConfig      `toml:"database"`
}

type TServiceConfig struct {
	Address  string
	LogLevel string
}

type TDbConfig struct {
	URI           string `toml:"uri" opt:"-"`
	DB            string `toml:"db" opt:"-"`
	Connect       string `toml:"connect" opt:"connect"`
	ReplicaSet    string `toml:"replica_set" opt:"replicaSet"`
	MaxPoolSize   int    `toml:"max_pool_size" opt:"maxPoolSize"`
	MinPoolSize   int    `toml:"min_pool_size" opt:"minPoolSize"`
	MaxIdleTimeMS int    `toml:"max_idle_time" opt:"maxIdleTimeMS"` // ms
	SSL           bool   `toml:"ssl" opt:"ssl"`
	SQLAddress    string
	SQLName       string
	SQLPassword   string
	SQLDbName     string
}

var Config TConfig

func InitConfig(configPath string) {
	if _, err := toml.DecodeFile(configPath, &Config); err != nil {
		panic(err)

	}
}
