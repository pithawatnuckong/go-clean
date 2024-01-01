package environment

import (
	"github.com/pithawatnuckong/go-clean/exception"
	"github.com/spf13/viper"
	"strings"
)

type DatabaseEnv struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	SslMode  string
}

type LoggingEnv struct {
	Level string
}

type Configuration struct {
	Database DatabaseEnv
	Logging  LoggingEnv
}

func NewEnvironment() (*Configuration, Environment) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("environment")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	exception.PanicLogging(viper.ReadInConfig())

	var configuration Configuration
	exception.PanicLogging(viper.Unmarshal(&configuration))

	return &configuration, &environmentImpl{}
}
