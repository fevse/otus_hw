package config

import (
	"github.com/BurntSushi/toml"
)

// При желании конфигурацию можно вынести в internal/config.
// Организация конфига в main принуждает нас сужать API компонентов, использовать
// при их конструировании только необходимые параметры, а также уменьшает вероятность циклической зависимости.
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
