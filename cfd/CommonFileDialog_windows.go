// +build windows

package cfd

import "github.com/go-ole/go-ole"

func Initialize() {
	// Swallow error
	_ = ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED|ole.COINIT_DISABLE_OLE1DDE)
}

func NewOpenFileDialog(config DialogConfig) (OpenFileDialog, error) {
	Initialize()

	openDialog, err := newIFileOpenDialog()
	if err != nil {
		return nil, err
	}
	err = config.apply(openDialog)
	if err != nil {
		return nil, err
	}
	return openDialog, nil
}

func NewOpenMultipleFilesDialog(config DialogConfig) (OpenMultipleFilesDialog, error) {
	Initialize()

	openDialog, err := newIFileOpenDialog()
	if err != nil {
		return nil, err
	}
	err = config.apply(openDialog)
	if err != nil {
		return nil, err
	}
	err = openDialog.setIsMultiselect(true)
	if err != nil {
		return nil, err
	}
	return openDialog, nil
}

func NewSelectFolderDialog(config DialogConfig) (SelectFolderDialog, error) {
	Initialize()

	openDialog, err := newIFileOpenDialog()
	if err != nil {
		return nil, err
	}
	err = config.apply(openDialog)
	if err != nil {
		return nil, err
	}
	err = openDialog.setPickFolders(true)
	if err != nil {
		return nil, err
	}
	return openDialog, nil
}

func NewSaveFileDialog(config DialogConfig) (SaveFileDialog, error) {
	Initialize()

	saveDialog, err := newIFileSaveDialog()
	if err != nil {
		return nil, err
	}
	err = config.apply(saveDialog)
	if err != nil {
		return nil, err
	}
	return saveDialog, nil
}
