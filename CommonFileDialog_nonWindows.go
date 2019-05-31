// +build !windows

package cfd

import "fmt"

var unsupportedError = fmt.Errorf("CFD is not supported on non-windows platforms")

func NewOpenFileDialog(dialogTitle, defaultFolder, fileFilter string) (Dialog, error) {
	return nil, unsupportedError
}

func NewPickFolderDialog(dialogTitle, defaultFolder string) (Dialog, error) {
	return nil, unsupportedError
}

func NewSaveFileDialog(dialogTitle, defaultFolder, fileFilter string) (Dialog, error) {
	return nil, unsupportedError
}
