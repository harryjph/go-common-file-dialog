// +build !windows

package cfdutil

import (
	"fmt"
)

var unsupportedError = fmt.Errorf("CFD is not supported on non-windows platforms")

func ShowOpenFileDialog(config DialogConfig) (string, error) {
	return "", unsupportedError
}

func ShowOpenMultipleFilesDialog(config DialogConfig) ([]string, error) {
	return "", unsupportedError
}

func ShowPickFolderDialog(config DialogConfig) (string, error) {
	return "", unsupportedError
}

func ShowSaveFileDialog(config DialogConfig) (string, error) {
	return "", unsupportedError
}
