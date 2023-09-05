package config

import "flag"

type Config struct {
	DBConnSpec  string
	ServiceHost string
	ServicePort string
}

var conf *Config
var configPath = flag.String("conf", "", "path of conf")

func Load(configPath string) error {
	//
	//
	conf = &Config{}
	return nil
}
func Read() Config {

	return *conf
}
