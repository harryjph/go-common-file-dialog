// +build !windows

package cfd

import "fmt"

var unsupportedError = fmt.Errorf("Common File Dialogs are not available on non-Windows platforms")

func NewOpenFileDialog(dialogTitle, defaultFolder, fileFilter string) (Dialog, error) {
	return nil, unsupportedError
}

func NewOpenMultipleFileDialog(config DialogConfig) (OpenMultipleDialog, error) {
	return nil, unsupportedError
}

func NewPickFolderDialog(dialogTitle, defaultFolder string) (Dialog, error) {
	return nil, unsupportedError
}

func NewSaveFileDialog(dialogTitle, defaultFolder, fileFilter string) (Dialog, error) {
	return nil, unsupportedError
}
