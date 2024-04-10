package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	// TODO
	Logger     LoggerConf
	HTTPServer HTTPconf
	DB         DBConf
}

type LoggerConf struct {
	Level string
	// TODO
}

type HTTPconf struct {
	Host string
	Port string
}

type DBConf struct {
	Type string
	DSN  string
}

func NewConfig(configFile string) (c Config, err error) {
	_, err = toml.DecodeFile(configFile, &c)
	if err != nil {
		return Config{}, err
	}
	return c, nil
}

// TODO
