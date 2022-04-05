package util

import (
	"errors"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"

	"github.com/Mopip77/screenshot-handler/config"
	"github.com/Mopip77/screenshot-handler/consts"

	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli/v2"
)

func LoadScreenshot(ctx *cli.Context) (imageName, imagePath string, imageContent []byte, fromClipboard bool, err error) {

	if ctx.Bool("from-dir") {
		imagePath, err = getLatestScreenshotFile()
		if err != nil {
			return
		}
	} else if ctx.String("file") != "" {
		imagePath = ctx.String("file")
		imagePath, err = homedir.Expand(imagePath)
		if err != nil {
			return
		}
		imagePath, err = filepath.Abs(imagePath)
		if err != nil {
			return
		}
	} else {
		fromClipboard = true
		imageName = consts.SCREENSHOT_FILENAME
		imageContent, err = ReadImageFromClipboard()
		if err != nil {
			return
		}
	}

	if !fromClipboard {
		splits := strings.Split(imagePath, "/")
		imageName = splits[len(splits)-1]
		imageContent, err = ioutil.ReadFile(imagePath)
		if err != nil {
			return
		}
	}

	return
}

func getLatestScreenshotFile() (string, error) {
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
