package bootstrap

type Config struct {
	AppName string
	Host    string
	Port    string
}

func NewConfig(appName, port string) *Config {
	return &Config{
		AppName: appName,
		Port:    port,
	}
}
