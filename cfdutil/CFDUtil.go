// +build windows

package cfdutil

import (
	"github.com/harry1453/go-common-file-dialog/cfd"
)

func ShowOpenFileDialog(config cfd.DialogConfig) (string, error) {
	openDialog, err := cfd.NewOpenFileDialog(config)
	if err != nil {
		return "", err
	}
	defer openDialog.Release()
	return openDialog.ShowAndGet()
}

func ShowOpenMultipleFilesDialog(config cfd.DialogConfig) ([]string, error) {
	openDialog, err := cfd.NewOpenMultipleFileDialog(config)
	if err != nil {
		return nil, err
	}
	defer openDialog.Release()
	return openDialog.ShowAndGetAll()
}

func ShowPickFolderDialog(config cfd.DialogConfig) (string, error) {
	openDialog, err := cfd.NewPickFolderDialog(config)
	if err != nil {
		return "", err
	}
	defer openDialog.Release()
	return openDialog.ShowAndGet()
}

func ShowSaveFileDialog(config cfd.DialogConfig) (string, error) {
	saveDialog, err := cfd.NewSaveFileDialog(config)
	if err != nil {
		return "", err
	}
	defer saveDialog.Release()
	return saveDialog.ShowAndGet()
}
