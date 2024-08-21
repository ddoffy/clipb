package handlers

import "github.com/atotto/clipboard"

func GetClipboardContent() (string, error) {
	content, err := clipboard.ReadAll()
	if err != nil {
		return "", err
	}

	return content, nil
}
