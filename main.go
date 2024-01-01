package main

import (
	"github.com/pithawatnuckong/go-clean/env"
)

func main() {
	config, finder := env.NewEnvironment()
	_, _ = config, finder
}
