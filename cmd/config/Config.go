package config

import (
	"fmt"
	"os"

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

//Load() - функция обращается к файлу конфигурации и записывает соответствующее объекту Config содержимое в структуру.
func Load(configPath string) error {
	if conf != nil {
		return nil
	}
	check_file, err := os.Stat(configPath)
	if err != nil {
		return fmt.Errorf("ошибка чтения структуры файла config, файл поврежден")
	}
	if check_file.Size() == 0 {
		return fmt.Errorf("файл конфигурации пуст или повреждён")
	}
	cfg, err := ini.Load(configPath)
	if err != nil {
		return fmt.Errorf("CONFIG - Ошибка чтения файла config: %s", err)
	}
	conf = &Config{
		DBConnSpec:  cfg.Section("").Key("DBConnSpec").String(),
		ServiceHost: cfg.Section("").Key("ServiceHost").String(),
		ServicePort: cfg.Section("").Key("ServicePort").String(),
	}
	return err
}

//Read() - функция реализует чтение данных из структуры Config из пакетной переменной conf.
func Read() Config {
	return *conf
}
