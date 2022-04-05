package util_test

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/Mopip77/screenshot-handler/config"
	"github.com/Mopip77/screenshot-handler/util"
)

func TestTencentOcr(t *testing.T) {
	config.InitConfig()

	imageContent, err := ioutil.ReadFile("../test/assets/test-ocr.png")
	if err != nil {
		t.Error(err)
	}

	if ocrResult, err := util.OcrTencent(imageContent, false, false); err != nil {
		t.Error(err)
	} else {
		fmt.Println("ocr result:", ocrResult)
	}

}
