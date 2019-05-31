package cfdutil

import "github.com/harry1453/go-common-file-dialog"

func ShowOpenFileDialog(dialogTitle, defaultFolder, typeFilter string) (string, error) {
	if err := cfd.Initialize(); err != nil {
		return "", nil
	}
	defer cfd.UnInitialize()

	openDialog, err := cfd.NewOpenFileDialog(dialogTitle, defaultFolder, typeFilter)
	if err != nil {
		return "", err
	}
	return openDialog.ShowAndGet()
}

func ShowOpenFolderDialog(dialogTitle, defaultFolder string) (string, error) {
	if err := cfd.Initialize(); err != nil {
		return "", nil
	}
	defer cfd.UnInitialize()

	openDialog, err := cfd.NewPickFolderDialog(dialogTitle, defaultFolder)
	if err != nil {
		return "", err
	}
	return openDialog.ShowAndGet()
}

func ShowSaveFileDialog(dialogTitle, defaultFolder, typeFilter string) (string, error) {
	if err := cfd.Initialize(); err != nil {
		return "", nil
	}
	defer cfd.UnInitialize()

	saveDialog, err := cfd.NewSaveFileDialog(dialogTitle, defaultFolder, typeFilter)
	if err != nil {
		return "", err
	}
	return saveDialog.ShowAndGet()
}
