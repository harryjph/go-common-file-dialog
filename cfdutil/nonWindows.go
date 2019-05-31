// +build !windows

package cfdutil

import (
	"fmt"
)

var unsupportedError = fmt.Errorf("CFD is not supported on non-windows platforms")

func ShowOpenFileDialog(dialogTitle, defaultFolder, typeFilter string) (string, error) {
	return "", unsupportedError
}

func ShowOpenFolderDialog(dialogTitle, defaultFolder string) (string, error) {
	return "", unsupportedError
}

func ShowSaveFileDialog(dialogTitle, defaultFolder, typeFilter string) (string, error) {
	return "", unsupportedError
}
