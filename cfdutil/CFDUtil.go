package cfdutil

import "github.com/harry1453/go-common-file-dialog"

func ShowOpenFileDialog(dialogTitle /*TODO*/, defaultFolder, typeFilter /*TODO*/ string) (string, error) {
	if err := cfd.Initialize(); err != nil {
		return "", nil
	}
	defer cfd.UnInitialize()

	if openDialog, err := cfd.NewOpenFileDialog(); err == nil {
		defer openDialog.Release()
		if defaultFolder != "" {
			if err := openDialog.SetDefaultFolder(defaultFolder); err != nil {
				return "", err
			}
		}
		if err := openDialog.Show(); err == nil {
			if result, err := openDialog.GetResult(); err == nil {
				return result, nil
			} else {
				return "", err
			}
		} else {
			return "", err
		}
	} else {
		return "", err
	}
}

func ShowOpenFolderDialog(dialogTitle /*TODO*/, defaultFolder string) (string, error) {
	if err := cfd.Initialize(); err != nil {
		return "", nil
	}
	defer cfd.UnInitialize()

	if openDialog, err := cfd.NewOpenFileDialog(); err == nil {
		defer openDialog.Release()
		if err := openDialog.SetPickFolders(true); err == nil {
			if defaultFolder != "" {
				if err := openDialog.SetDefaultFolder(defaultFolder); err != nil {
					return "", err
				}
			}
			if err := openDialog.Show(); err == nil {
				if result, err := openDialog.GetResult(); err == nil {
					return result, nil
				} else {
					return "", err
				}
			} else {
				return "", err
			}
		} else {
			return "", err
		}
	} else {
		return "", err
	}
}

func ShowSaveFileDialog(dialogTitle /*TODO*/, defaultFolder, typeFilter /*TODO*/ string) (string, error) {
	if err := cfd.Initialize(); err != nil {
		return "", nil
	}
	defer cfd.UnInitialize()

	if saveDialog, err := cfd.NewSaveFileDialog(); err == nil {
		defer saveDialog.Release()
		if defaultFolder != "" {
			if err := saveDialog.SetDefaultFolder(defaultFolder); err != nil {
				return "", err
			}
		}
		if err := saveDialog.Show(); err == nil {
			if result, err := saveDialog.GetResult(); err == nil {
				return result, nil
			} else {
				return "", err
			}
		} else {
			return "", err
		}
	} else {
		return "", err
	}
}
