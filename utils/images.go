package utils

import (
	"errors"
	"strings"
)

func GetImage(text string) (string, error) {
	if len(text) <= 0 {
		return "", errors.New("empty arg")
	}
	image := SendRequest(strings.ReplaceAll(text, " ", ""))
	return image, nil
}
