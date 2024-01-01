package env

import (
	"github.com/spf13/viper"
	"os"
)

type Environment interface {
	Get(key string) string
}

type environmentImpl struct{}

func (config *environmentImpl) Get(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return viper.GetString(key)
}
