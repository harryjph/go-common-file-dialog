// +build !windows

package cfd

import "fmt"

var unsupportedError = fmt.Errorf("common file dialogs are only available on windows")

func NewOpenFileDialog(config DialogConfig) (OpenFileDialog, error) {
	return nil, unsupportedError
}

func NewOpenMultipleFilesDialog(config DialogConfig) (OpenMultipleFilesDialog, error) {
	return nil, unsupportedError
}

func NewSelectFolderDialog(config DialogConfig) (SelectFolderDialog, error) {
	return nil, unsupportedError
}

func NewSaveFileDialog(config DialogConfig) (SaveFileDialog, error) {
	return nil, unsupportedError
}
