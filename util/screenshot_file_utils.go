package util

import (
	"errors"
	"io/ioutil"
	"path/filepath"
	"sort"

	"github.com/Mopip77/screenshot-handler/config"

	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli/v2"
)

func GetLatestScreenshotPath(ctx *cli.Context) (string, error) {
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
	imagePath := filepath.Join(folder, files[0].Name())

	return imagePath, nil
}
