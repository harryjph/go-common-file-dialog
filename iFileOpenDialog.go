// +build windows

package cfd

import (
	"github.com/go-ole/go-ole/oleutil"
	"github.com/harry1453/go-common-file-dialog/util"
	"unsafe"
)

type iFileOpenDialog struct {
	vtbl *iFileOpenDialogVtbl
}

type iFileOpenDialogVtbl struct {
	iFileDialogVtbl

	GetResults       uintptr
	GetSelectedItems uintptr
}

const (
	clsidFileopendialog = "{DC1C5A9C-E88A-4dde-A5A1-60F82A20AEF7}"
)

func newIFileOpenDialog() (*iFileOpenDialog, error) {
	if unknown, err := oleutil.CreateObject(clsidFileopendialog); err == nil {
		return (*iFileOpenDialog)(unsafe.Pointer(unknown)), nil
	} else {
		return nil, err
	}
}

func (fileOpenDialog *iFileOpenDialog) Show() error {
	return fileOpenDialog.vtbl.show(unsafe.Pointer(fileOpenDialog))
}

func (fileOpenDialog *iFileOpenDialog) ShowAndGet() (string, error) {
	if err := fileOpenDialog.Show(); err != nil {
		return "", err
	}
	return fileOpenDialog.GetResult()
}

func (fileOpenDialog *iFileOpenDialog) Close() error {
	return fileOpenDialog.vtbl.close(unsafe.Pointer(fileOpenDialog))
}

func (fileOpenDialog *iFileOpenDialog) SetTitle(title string) error {
	return fileOpenDialog.vtbl.setTitle(unsafe.Pointer(fileOpenDialog), title)
}

func (fileOpenDialog *iFileOpenDialog) GetResult() (string, error) {
	return fileOpenDialog.vtbl.getResultString(unsafe.Pointer(fileOpenDialog))
}

func (fileOpenDialog *iFileOpenDialog) Release() error {
	return fileOpenDialog.vtbl.release(unsafe.Pointer(fileOpenDialog))
}

func (fileOpenDialog *iFileOpenDialog) SetDefaultFolder(defaultFolderPath string) error {
	return fileOpenDialog.vtbl.setDefaultFolder(unsafe.Pointer(fileOpenDialog), defaultFolderPath)
}

func (fileOpenDialog *iFileOpenDialog) SetInitialFolder(defaultFolderPath string) error {
	return fileOpenDialog.vtbl.setFolder(unsafe.Pointer(fileOpenDialog), defaultFolderPath)
}

func (fileOpenDialog *iFileOpenDialog) SetFileFilter(filter []FileFilter) error {
	return fileOpenDialog.vtbl.setFileTypes(unsafe.Pointer(fileOpenDialog), filter)
}

func (fileOpenDialog *iFileOpenDialog) SetRole(role string) error {
	return fileOpenDialog.vtbl.setClientGuid(unsafe.Pointer(fileOpenDialog), util.StringToUUID(role))
}

func (fileOpenDialog *iFileOpenDialog) setPickFolders(pickFolders bool) error {
	const FosPickfolders = 0x20
	if pickFolders {
		return fileOpenDialog.vtbl.addOption(unsafe.Pointer(fileOpenDialog), FosPickfolders)
	} else {
		return fileOpenDialog.vtbl.removeOption(unsafe.Pointer(fileOpenDialog), FosPickfolders)
	}
}
