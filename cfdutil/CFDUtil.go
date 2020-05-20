// +build windows

package cfdutil

import (
	"github.com/harry1453/go-common-file-dialog/cfd"
)

func ShowOpenFileDialog(config cfd.DialogConfig) (string, error) {
	dialog, err := cfd.NewOpenFileDialog(config)
	if err != nil {
		return "", err
	}
	defer dialog.Release()
	return dialog.ShowAndGetResult()
}

func ShowOpenMultipleFilesDialog(config cfd.DialogConfig) ([]string, error) {
	dialog, err := cfd.NewOpenMultipleFilesDialog(config)
	if err != nil {
		return nil, err
	}
	defer dialog.Release()
	return dialog.ShowAndGetResults()
}

func ShowPickFolderDialog(config cfd.DialogConfig) (string, error) {
	dialog, err := cfd.NewSelectFolderDialog(config)
	if err != nil {
		return "", err
	}
	defer dialog.Release()
	return dialog.ShowAndGetResult()
}

func ShowSaveFileDialog(config cfd.DialogConfig) (string, error) {
	dialog, err := cfd.NewSaveFileDialog(config)
	if err != nil {
		return "", err
	}
	defer dialog.Release()
	return dialog.ShowAndGetResult()
}
