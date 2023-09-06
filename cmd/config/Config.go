package config

import "flag"

// Объект, реализующий чтение файла конфигурации.
// Файл содержит параметры для соединения с базой данных и параметры для соединения с веб-сервисом.
type Config struct {
	DBConnSpec  string
	ServiceHost string
	ServicePort string
}

var conf *Config
var configPath = flag.String("conf", "", "path of conf")

//Load() - функция обращается к файлу и записывает соответствующее объекту Config содержимое в структуру.
func Load(configPath string) error {
	//
	//
	conf = &Config{}
	return nil
}

//Read() - функция реализует чтение данных из структуры Config.
func Read() Config {
	return *conf
}
