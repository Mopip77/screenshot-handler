package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Mopip77/screenshot-handler/consts"
	"github.com/Mopip77/screenshot-handler/infra/output"
	"github.com/fatih/color"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

type Config struct {
	ScreenshotFolder string `mapstructure:"screenshot_folder"`

	Upload struct {
		Use       []string `mapstructure:"use"`
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

func InitConfig() (err error) {
	if home, err := os.UserHomeDir(); err != nil {
		panic(err)
	} else {
		_ConfigFilePath = fmt.Sprintf("%s/%s", home, consts.CONFIG_FILE_NAME)
	}

	config.WithOptions(config.ParseEnv)
	config.AddDriver(yaml.Driver)
	if _, err = os.Stat(_ConfigFilePath); errors.Is(err, os.ErrNotExist) {
		if err = initTemplate(); err != nil {
			return err
		}
	}
	err = config.LoadFiles(_ConfigFilePath)
	if err != nil {
		return
	}

	config.BindStruct("", &GlobalConfig)
	return nil
}

func initTemplate() error {
	output.Fmt.Printf("config file %s not found, ", _ConfigFilePath)
	output.CyanFmt.Add(color.Bold).Println("generate default config.")

	return createDefaultConfigFileTemplate()
}

func createDefaultConfigFileTemplate() error {
	defaultConfigFileTemplate := `
# default screenshot folder
screenshot_folder: <change-to-your-screenshot-folder>

# upload to image host settings
upload:
  use: # smms | github
  smms_token: 
  github:
    username: 
    repo: 
    token: 

# ocr settings
ocr:
  use: # tencent
  tencent:
    secret_id: 
    secret_key: 
`

	if err := ioutil.WriteFile(_ConfigFilePath, []byte(defaultConfigFileTemplate), 0644); err != nil {
		return err
	}

	return nil
}
