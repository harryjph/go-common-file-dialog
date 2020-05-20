// +build !windows

package cfd

import "fmt"

var unsupportedError = fmt.Errorf("common file dialogs are only available on windows")

func NewOpenFileDialog(config DialogConfig) (Dialog, error) {
	return nil, unsupportedError
}

func NewOpenMultipleFilesDialog(config DialogConfig) (OpenMultipleDialog, error) {
	return nil, unsupportedError
}

func NewPickFolderDialog(config DialogConfig) (Dialog, error) {
	return nil, unsupportedError
}

func NewSaveFileDialog(config DialogConfig) (Dialog, error) {
	return nil, unsupportedError
}
