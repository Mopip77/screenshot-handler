package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Mopip77/screenshot-handler/infra/output"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

type Config struct {
	ScreenshotFolder string `mapstructure:"screenshot_folder"`

	Upload struct {
		Use       string `mapstructure:"use"`
		SmmsToken string `mapstructure:"smms_token"`
		Github    struct {
			Username string `mapstructure:"username"`
			Repo     string `mapstructure:"repo"`
			Token    string `mapstructure:"token"`
		} `mapstructure:"github"`
	} `mapstructure:"upload"`

	Ocr struct {
		Use     string `mapstructure:"use"`
		Tencent struct {
			SecretId  string `mapstructure:"secret_id"`
			SecretKey string `mapstructure:"secret_key"`
		} `mapstructure:"tencent"`
	} `mapstructure:"ocr"`
}

var (
	GlobalConfig    *Config
	_ConfigFilePath string
)

const (
	_CONFIG_FILENAME = ".schrc.yaml"
)

func InitConfig() (err error) {
	if home, err := os.UserHomeDir(); err != nil {
		panic(err)
	} else {
		_ConfigFilePath = fmt.Sprintf("%s/%s", home, _CONFIG_FILENAME)
	}

	config.WithOptions(config.ParseEnv)
	config.AddDriver(yaml.Driver)
	if _, err = os.Stat(_ConfigFilePath); errors.Is(err, os.ErrNotExist) {
		output.RedFmt.Printf("config file %s not found\n", _ConfigFilePath)
		createDefaultConfigFileTemplate()
		output.RedFmt.Printf("default config file %s created, now set screenshot folder in %s\n", _ConfigFilePath, _ConfigFilePath)
		return
	}
	err = config.LoadFiles(_ConfigFilePath)
	if err != nil {
		return
	}

	config.BindStruct("", &GlobalConfig)
	return nil
}

func createDefaultConfigFileTemplate() {
	defaultConfigFileTemplate := `screenshot_folder: <change-to-your-screenshot-folder>`

	if err := ioutil.WriteFile(_ConfigFilePath, []byte(defaultConfigFileTemplate), 0644); err != nil {
		output.Fmt.Printf("generate default config file %s failed, %s", _ConfigFilePath, err)
	}
}
