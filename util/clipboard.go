package util

import (
	"fmt"

	"golang.design/x/clipboard"
)

type ClipboardFormatTypeEnum int

const (
	CLIPBOARD_FORMAT_TEXT = iota
	CLIPBOARD_FORMAT_IMAGE
)

func ReadImageFromClipboard() ([]byte, error) {
	if err := clipboard.Init(); err != nil {
		return nil, err
	}

	imageContent := clipboard.Read(clipboard.FmtImage)
	if imageContent == nil {
		return nil, fmt.Errorf("no image found in clipboard")
	}

	return imageContent, nil
}


func WriteToClipboard(formatType ClipboardFormatTypeEnum, content []byte) error {
	if err := clipboard.Init(); err != nil {
		return err
	}

	switch formatType {
	case CLIPBOARD_FORMAT_TEXT:
		clipboard.Write(clipboard.FmtText, content)
	case CLIPBOARD_FORMAT_IMAGE:
		clipboard.Write(clipboard.FmtImage, content)
	}

	return nil
}
