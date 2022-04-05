package config_test

import (
	"fmt"
	"testing"

	"github.com/Mopip77/screenshot-handler/config"
)

func TestLoad(*testing.T) {
	if err := config.InitConfig(); err != nil {
		panic(err)
	}

	fmt.Println("output:", config.GlobalConfig.Upload.Use)
}
