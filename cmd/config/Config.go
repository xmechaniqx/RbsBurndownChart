package config

import (
	"fmt"

	"gopkg.in/ini.v1"
)

// Объект, реализующий чтение файла конфигурации.
// Файл содержит параметры для соединения с базой данных и параметры для соединения с веб-сервисом.
type Config struct {
	DBConnSpec  string
	ServiceHost string
	ServicePort string
}

var conf *Config

//Load() - функция обращается к файлу и записывает соответствующее объекту Config содержимое в структуру.
func Load(configPath string) error {
	if conf != nil {
		return nil
	}
	cfg, err := ini.Load(configPath)
	if err != nil {
		fmt.Println("CONFIG - Ошибка чтения файла config или пути configPath", err)
	}
	conf = &Config{
		DBConnSpec:  cfg.Section("").Key("DBConnSpec").String(),
		ServiceHost: cfg.Section("").Key("ServiceHost").String(),
		ServicePort: cfg.Section("").Key("ServicePort").String(),
	}
	return nil
}

//Read() - функция реализует чтение данных из структуры Config из пакетной переменной conf.
func Read() Config {
	return *conf
}
