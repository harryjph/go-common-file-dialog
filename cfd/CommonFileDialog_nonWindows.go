// +build !windows

package cfd

import "fmt"

var unsupportedError = fmt.Errorf("Common File Dialogs are not available on non-Windows platforms")

func NewOpenFileDialog(config DialogConfig) (Dialog, error) {
	return nil, unsupportedError
}

func NewOpenMultipleFileDialog(config DialogConfig) (OpenMultipleDialog, error) {
	return nil, unsupportedError
}

func NewPickFolderDialog(config DialogConfig) (Dialog, error) {
	return nil, unsupportedError
}

func NewSaveFileDialog(config DialogConfig) (Dialog, error) {
	return nil, unsupportedError
}
