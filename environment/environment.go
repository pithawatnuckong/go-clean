package environment

import (
	"github.com/spf13/viper"
)

type Environment interface {
	Get(key string) string
}

type environmentImpl struct{}

func (config *environmentImpl) Get(key string) string {
	return viper.GetString(key)
}
