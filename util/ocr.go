package util

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/Mopip77/screenshot-handler/config"

	"github.com/fatih/color"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	ocr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ocr/v20181119"
)

// ---------------------------- Tencent ---------------------------------------

type OcrTencentResponse struct {
	Response struct {
		TextDetctions []struct {
			DetectedText string `json:"DetectedText"`
			ItemPolygon  struct {
				// left-top x, y, width, height after image rotation
				X      int `json:"X"`
				Y      int `json:"Y"`
				Width  int   `json:"Width"`
				Height int   `json:"Height"`
			} `json:"ItemPolygon"`
		} `json:"TextDetections"`
	} `json:"Response"`
}

func OcrTencent(imageContent []byte, linefeed, useFullWidth bool) (string, error) {
	credential := common.NewCredential(
		config.GlobalConfig.Ocr.Tencent.SecretId,
		config.GlobalConfig.Ocr.Tencent.SecretKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "ocr.tencentcloudapi.com"
	client, _ := ocr.NewClient(credential, "ap-guangzhou", cpf)

	request := ocr.NewGeneralBasicOCRRequest()
	imageBase64 := base64.StdEncoding.EncodeToString(imageContent)
	request.ImageBase64 = &imageBase64

	resp, err := client.GeneralBasicOCR(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok || err != nil {
		return "", err
	}

	var response OcrTencentResponse
	if err := json.Unmarshal([]byte(resp.ToJsonString()), &response); err != nil {
		return "", err
	}

	var seperator string
	if linefeed {
		seperator = "\n"
	} else {
		seperator = " "
	}

	var ocrResult string
	texts := response.Response.TextDetctions
	for idx, v := range texts {
		if useFullWidth {
			v.DetectedText = half2Full(v.DetectedText)
		}

		if idx == 0 || texts[idx-1].ItemPolygon.Y + texts[idx-1].ItemPolygon.Height > v.ItemPolygon.Y {
			ocrResult += v.DetectedText
			continue
		}
		
		ocrResult += seperator + v.DetectedText
	}

	return ocrResult, nil
}

func CheckRequiredOcrTencentConfig() error {
	if config.GlobalConfig.Ocr.Tencent.SecretId == "" {
		return fmt.Errorf(color.BlueString("ocr.tencent.secret_id") + " is not set, please set it in config file")
	}
	if config.GlobalConfig.Ocr.Tencent.SecretKey == "" {
		return fmt.Errorf(color.BlueString("ocr.tencent.secret_key") + " is not set, please set it in config file")
	}
	return nil
}

func half2Full(str string) string {
	var dict = map[string]string{
		",": "，", ".": "。", "!": "！", "?": "？", "(": "（", ")": "）", "[": "【", "]": "】", ":": "：",
	}

	var result string
	for _, c := range str {
		if dict[string(c)] != "" {
			result += dict[string(c)]
		} else {
			result += string(c)
		}
	}
	return result
}
