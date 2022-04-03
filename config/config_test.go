package config

import "testing"

func TestLoad(*testing.T) {
	if err := InitConfig(); err != nil {
		panic(err)
	}
}
