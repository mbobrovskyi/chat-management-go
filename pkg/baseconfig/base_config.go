package baseconfig

type BaseConfig struct {
	Environment    Environment `env:"ENVIRONMENT" envDefault:"development"`
	HttpServerAddr string      `env:"HTTP_SERVER_ADDR" envDefault:"0.0.0.0:8081"`
	LogLevel       string      `env:"LOG_LEVEL" envDefault:"debug"`
}
