package environment

import (
	"github.com/pithawatnuckong/go-clean/exception"
	"github.com/spf13/viper"
	"os"
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

type RedisEnv struct {
	Host     string
	Password string
	Port     int
}

type Configuration struct {
	Database DatabaseEnv
	Logging  LoggingEnv
	Redis    RedisEnv
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

	// Bind OS environment for viper
	exception.PanicLogging(viper.BindEnv(os.Environ()...))

	return &configuration, &environmentImpl{}
}
