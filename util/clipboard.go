package util

import "golang.design/x/clipboard"

type ClipboardFormatTypeEnum int

const (
	CLIPBOARD_FORMAT_TEXT = iota
	CLIPBOARD_FORMAT_IMAGE
)

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
