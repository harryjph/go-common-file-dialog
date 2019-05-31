// +build windows

package cfdutil

import . "github.com/harry1453/go-common-file-dialog"

func ShowOpenFileDialog(config DialogConfig) (string, error) {
	if err := Initialize(); err != nil {
		return "", nil
	}
	defer UnInitialize()

	openDialog, err := NewOpenFileDialog(config)
	if err != nil {
		return "", err
	}
	return openDialog.ShowAndGet()
}

func ShowPickFolderDialog(config DialogConfig) (string, error) {
	if err := Initialize(); err != nil {
		return "", nil
	}
	defer UnInitialize()

	openDialog, err := NewPickFolderDialog(config)
	if err != nil {
		return "", err
	}
	return openDialog.ShowAndGet()
}

func ShowSaveFileDialog(config DialogConfig) (string, error) {
	if err := Initialize(); err != nil {
		return "", nil
	}
	defer UnInitialize()

	saveDialog, err := NewSaveFileDialog(config)
	if err != nil {
		return "", err
	}
	return saveDialog.ShowAndGet()
}
