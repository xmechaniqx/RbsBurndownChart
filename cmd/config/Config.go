package config

type Config struct {
	DBConnSpec  string
	ServiceHost string
	ServicePort string
}

func Load(configPath string) error {
	return nil
}
func Read() Config {

	return Config{}
}
