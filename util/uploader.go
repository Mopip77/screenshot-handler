package util

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"screenshot-handler/config"
	"screenshot-handler/consts"
	"time"

	"github.com/fatih/color"
	"github.com/go-resty/resty/v2"
)

// ---------------------------- SMMS ----------------------------------------

// api docs: https://doc.sm.ms/#api-Image-Upload

// response example:
// {"success":true,"code":"success","message":"Upload success.","data":{"file_id":0,"width":220,"height":220,"filename":"image.png","storename":"edQp6ViCKMoDcEs.png","size":1785,"path":"\/2022\/04\/03\/edQp6ViCKMoDcEs.png","hash":"pxvdO5RSCJNKaf4TEncblV7Fie","url":"https:\/\/s2.loli.net\/2022\/04\/03\/edQp6ViCKMoDcEs.png","delete":"https:\/\/sm.ms\/delete\/pxvdO5RSCJNKaf4TEncblV7Fie","page":"https:\/\/sm.ms\/image\/edQp6ViCKMoDcEs"},"RequestId":"514D3848-FC59-4861-80A1-5FBB921345E7"}

type SmmsUploadResponse struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Data    SmmsUploadData `json:"data"`
}

type SmmsUploadData struct {
	Url      string `json:"url"`
	ViewPage string `json:"page"`
}

func UploadToSmms(imageContent []byte) (string, error) {
	url := "https://sm.ms/api/v2/upload"

	resp, err := resty.New().R().
		SetHeader("Authorization", config.GlobalConfig.Upload.SmmsToken).
		SetHeader("Content-Type", "multipart/form-data").
		SetFileReader("smfile", "image.png", bytes.NewReader(imageContent)).
		Post(url)
	if err != nil {
		return "", err
	}

	var response SmmsUploadResponse
	if err := json.Unmarshal(resp.Body(), &response); err != nil {
		return "", err
	}
	if !response.Success {
		return "", fmt.Errorf("upload to sm.ms failed, %s", response.Message)
	}

	return response.Data.Url, nil
}

func CheckRequiredSmmsConfig() error {
	if config.GlobalConfig.Upload.SmmsToken == "" {
		return fmt.Errorf("upload.smms_token is empty, please set it in config file")
	}
	return nil
}

// ---------------------------- GITHUB ----------------------------------------

// api docs: https://docs.github.com/cn/rest/reference/repos#create-or-update-file-contents

type GithubUploadResponse struct {
	Content struct {
		DownloadUrl string `json:"download_url"`
	} `json:"content"`
}

func UploadToGithub(imageContent []byte, useJsDeliver bool) (string, error) {
	path := fmt.Sprintf("%s/%s.png", time.Now().Format("20060102"), time.Now().Format("150405"))

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s",
		config.GlobalConfig.Upload.Github.Username,
		config.GlobalConfig.Upload.Github.Repo,
		path)

	resp, err := resty.New().R().
		SetHeader("Accept", "application/vnd.github.v3+json").
		SetAuthToken(config.GlobalConfig.Upload.Github.Token).
		SetBody(map[string]string{
			"message": fmt.Sprintf("upload by screenshot-handler(%s)", consts.GITHUB_REPO),
			"content": base64.StdEncoding.EncodeToString(imageContent),
		}).
		Put(url)
	if err != nil {
		return "", err
	}
	if !resp.IsSuccess() {
		var errMsg string
		switch resp.StatusCode() {
		case 404:
			errMsg = "repo not found"
		case 409:
			errMsg = "confilct"
		case 422:
			errMsg = "validation failed"
		}
		return "", fmt.Errorf("upload to github failed, %s", errMsg)
	}

	var imageUrl string
	var response GithubUploadResponse
	if err := json.Unmarshal(resp.Body(), &response); err != nil {
		return "", err
	}
	imageUrl = response.Content.DownloadUrl

	if useJsDeliver {
		imageUrl = fmt.Sprintf("https://cdn.jsdelivr.net/gh/%s/%s/%s",
			config.GlobalConfig.Upload.Github.Username,
			config.GlobalConfig.Upload.Github.Repo,
			path)
	}

	return imageUrl, nil
}

func CheckRequiredGithubConfig() error {
	if config.GlobalConfig.Upload.Github.Username == "" {
		return fmt.Errorf("%s is empty, please set it in config file", color.BlueString("upload.github.username"))
	}
	if config.GlobalConfig.Upload.Github.Repo == "" {
		return fmt.Errorf("%s is empty, please set it in config file", color.BlueString("upload.github.repo"))
	}
	if config.GlobalConfig.Upload.Github.Token == "" {
		return fmt.Errorf("%s is empty, please set it in config file", color.BlueString("upload.github.token"))
	}
	return nil
}
