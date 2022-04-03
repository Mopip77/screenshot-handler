package config

import (
	"fmt"
	"testing"
)

func TestLoad(*testing.T) {
	if err := InitConfig(); err != nil {
		panic(err)
	}

	fmt.Println("output:", GlobalConfig.Upload.Use)
}
