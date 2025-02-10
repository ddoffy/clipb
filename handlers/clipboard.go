package handlers

import (
	"bytes"
	"fmt"

	"github.com/atotto/clipboard"
)

func GetClipboardContent() (string, error) {
	content, err := clipboard.ReadAll()
	if err != nil {
		return "", err
	}

	return content, nil
}

func GetImg() (bytes.Buffer, error) {
	// get clipboard content
	img := robotgo.ReadAll()

	if img == nil {
		return nil, fmt.Errorf("Error getting clipboard content")
	}

	// convert data  to binary
	var buffer bytes.Buffer
	err := robotgo.SaveBitmap(img, &buffer)

	if err != nil {
		return nil, fmt.Errorf("Error converting clipboard content to binary: %v", err)
	}

	return buffer, nil
}
