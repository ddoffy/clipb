package handlers

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/atotto/clipboard"
	"github.com/go-vgo/robotgo"
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

func GetImageFromClipboard() (bytes.Buffer, error) {
	// Use wayclip to capture the clipboard image (if available)
	cmd := exec.Command("wayclip", "get", "image")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("Error getting image from clipboard: %v", err)
	}

	// convert data to binary
	var buffer bytes.Buffer
	buffer.Write(output)

	return buffer, nil
}
