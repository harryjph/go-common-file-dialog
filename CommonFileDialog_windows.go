// +build windows

package cfd

func NewOpenFileDialog(dialogTitle, defaultFolder, fileFilter string) (Dialog, error) {
	if !initialized {
		if err := Initialize(); err != nil {
			return nil, err
		}
	}

	openDialog, err := newIFileOpenDialog()
	if err != nil {
		return nil, err
	}
	err = openDialog.SetTitle(dialogTitle)
	if err != nil {
		return nil, err
	}
	err = openDialog.SetDefaultFolder(defaultFolder)
	if err != nil {
		return nil, err
	}
	err = openDialog.setPickFolders(false)
	if err != nil {
		return nil, err
	}
	err = openDialog.SetFileFilter(fileFilter)
	if err != nil {
		return nil, err
	}
	return openDialog, nil
}

func NewPickFolderDialog(dialogTitle, defaultFolder string) (Dialog, error) {
	if !initialized {
		if err := Initialize(); err != nil {
			return nil, err
		}
	}

	openDialog, err := newIFileOpenDialog()
	if err != nil {
		return nil, err
	}
	err = openDialog.SetTitle(dialogTitle)
	if err != nil {
		return nil, err
	}
	err = openDialog.SetDefaultFolder(defaultFolder)
	if err != nil {
		return nil, err
	}
	err = openDialog.setPickFolders(true)
	if err != nil {
		return nil, err
	}
	return openDialog, nil
}

func NewSaveFileDialog(dialogTitle, defaultFolder, fileFilter string) (Dialog, error) {
	if !initialized {
		if err := Initialize(); err != nil {
			return nil, err
		}
	}

	saveDialog, err := newIFileSaveDialog()
	if err != nil {
		return nil, err
	}
	err = saveDialog.SetTitle(dialogTitle)
	if err != nil {
		return nil, err
	}
	err = saveDialog.SetDefaultFolder(defaultFolder)
	if err != nil {
		return nil, err
	}
	err = saveDialog.SetFileFilter(fileFilter)
	if err != nil {
		return nil, err
	}
	return saveDialog, nil
}
