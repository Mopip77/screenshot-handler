package util

import (
	"fmt"
	"io/ioutil"
	"screenshot-handler/config"
	"testing"
)

func TestTencentOcr(t *testing.T) {
	config.InitConfig()

	imageContent, err := ioutil.ReadFile("../test/assets/test-ocr.png")
	if err != nil {
		t.Error(err)
	}

	if ocrResult, err := OcrTencent(imageContent, false, false); err != nil {
		t.Error(err)
	} else {
		fmt.Println("ocr result:", ocrResult)
	}

}
