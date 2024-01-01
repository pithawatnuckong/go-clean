package env

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

type Configuration struct {
	Database DatabaseEnv
}

func NewEnvironment() (*Configuration, Environment) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("env")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	exception.PanicLogging(viper.ReadInConfig())

	var configuration Configuration
	exception.PanicLogging(viper.Unmarshal(&configuration))

	return &configuration, &environmentImpl{}
}
