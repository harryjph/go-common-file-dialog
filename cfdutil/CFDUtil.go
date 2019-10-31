// +build windows

package cfdutil

import . "github.com/harry1453/go-common-file-dialog"

func ShowOpenFileDialog(config DialogConfig) (string, error) {
	if err := Initialize(); err != nil {
		return "", err
	}
	defer UnInitialize()

	openDialog, err := NewOpenFileDialog(config)
	if err != nil {
		return "", err
	}
	defer openDialog.Release()
	return openDialog.ShowAndGet()
}

func ShowOpenMultipleFilesDialog(config DialogConfig) ([]string, error) {
	if err := Initialize(); err != nil {
		return nil, err
	}
	defer UnInitialize()

	openDialog, err := NewOpenMultipleFileDialog(config)
	if err != nil {
		return nil, err
	}
	defer openDialog.Release()
	return openDialog.ShowAndGetAll()
}

func ShowPickFolderDialog(config DialogConfig) (string, error) {
	if err := Initialize(); err != nil {
		return "", err
	}
	defer UnInitialize()

	openDialog, err := NewPickFolderDialog(config)
	if err != nil {
		return "", err
	}
	defer openDialog.Release()
	return openDialog.ShowAndGet()
}

func ShowSaveFileDialog(config DialogConfig) (string, error) {
	if err := Initialize(); err != nil {
		return "", err
	}
	defer UnInitialize()

	saveDialog, err := NewSaveFileDialog(config)
	if err != nil {
		return "", err
	}
	defer saveDialog.Release()
	return saveDialog.ShowAndGet()
}
