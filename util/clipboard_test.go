package util_test

import (
	"testing"

	"github.com/Mopip77/screenshot-handler/infra/output"
	"github.com/Mopip77/screenshot-handler/util"
)

func TestRead(t *testing.T) {
	content, err := util.ReadImageFromClipboard()
	if err != nil {
		t.Error(err)
	}

	output.GreenFmt.Println("result,", string(content))
}