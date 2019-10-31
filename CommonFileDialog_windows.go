// +build windows

package cfd

func NewOpenFileDialog(config DialogConfig) (Dialog, error) {
	if !initialized {
		if err := Initialize(); err != nil {
			return nil, err
		}
	}

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

func NewOpenMultipleFileDialog(config DialogConfig) (OpenMultipleDialog, error) {
	if !initialized {
		if err := Initialize(); err != nil {
			return nil, err
		}
	}

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

func NewPickFolderDialog(config DialogConfig) (Dialog, error) {
	if !initialized {
		if err := Initialize(); err != nil {
			return nil, err
		}
	}

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

func NewSaveFileDialog(config DialogConfig) (Dialog, error) {
	if !initialized {
		if err := Initialize(); err != nil {
			return nil, err
		}
	}

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
