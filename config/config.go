package config

import (
	"github.com/BurntSushi/toml"
)

type TConfig struct {
	ServiceConfig TServiceConfig `toml:"server"`
	DbConfig      TDbConfig      `toml:"database"`
}

type TServiceConfig struct {
	Address         string
	LogLevel        string
	ConvertServer   string
	CameraServer    string
	ExpireTokenHour int
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
	SqlAddress    string
	Name          string
	Password      string
	DbName        string
}

var Config TConfig

func InitConfig(configPath string) {
	if _, err := toml.DecodeFile(configPath, &Config); err != nil {
		panic(err)

	}
}
