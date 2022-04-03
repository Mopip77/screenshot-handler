package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"screenshot-handler/config"

	"github.com/go-resty/resty/v2"
)

// ---------------------------- SMMS ----------------------------------------

// api docs: https://doc.sm.ms/#api-Image-Upload

// response example:
// {"success":true,"code":"success","message":"Upload success.","data":{"file_id":0,"width":220,"height":220,"filename":"image.png","storename":"edQp6ViCKMoDcEs.png","size":1785,"path":"\/2022\/04\/03\/edQp6ViCKMoDcEs.png","hash":"pxvdO5RSCJNKaf4TEncblV7Fie","url":"https:\/\/s2.loli.net\/2022\/04\/03\/edQp6ViCKMoDcEs.png","delete":"https:\/\/sm.ms\/delete\/pxvdO5RSCJNKaf4TEncblV7Fie","page":"https:\/\/sm.ms\/image\/edQp6ViCKMoDcEs"},"RequestId":"514D3848-FC59-4861-80A1-5FBB921345E7"}

type SmmsUploadResponse struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Data SmmsUploadData `json:"data"`
}

type SmmsUploadData struct {
	Url string `json:"url"`
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
