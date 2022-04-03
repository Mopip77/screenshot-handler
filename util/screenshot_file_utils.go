package util

import (
	"errors"
	"io/ioutil"
	"path/filepath"
	"screenshot-handler/config"
	"sort"

	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli/v2"
)

func GetLatestScreenshotPath(ctx *cli.Context) (string, error) {
	imagePath := ctx.Args().First()
	if imagePath == "" {
		folder, err := homedir.Expand(config.GlobalConfig.ScreenshotFolder)
		if err != nil {
			return "", err
		}
		folder, err = filepath.Abs(folder)
		if err != nil {
			return "", err
		}
		files, err := ioutil.ReadDir(folder)
		if err != nil {
			return "", err
		}
		if len(files) == 0 {
			return "", errors.New("no screenshot file in " + folder)
		}

		sort.SliceStable(files, func(i, j int) bool {
			return files[i].ModTime().After(files[j].ModTime())
		})
		imagePath = filepath.Join(folder, files[0].Name())
	}
	imagePath, err := filepath.Abs(imagePath)
	if err != nil {
		return "", err
	}

	return imagePath, nil
}
