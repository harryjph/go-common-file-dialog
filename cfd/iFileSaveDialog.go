// +build windows

package cfd

import (
	"github.com/go-ole/go-ole"
	"github.com/harry1453/go-common-file-dialog/util"
	"unsafe"
)

var (
	saveFileDialogCLSID = ole.NewGUID("{C0B4E2F3-BA21-4773-8DBA-335EC946EB8B}")
	saveFileDialogIID   = ole.NewGUID("{84bccd23-5fde-4cdb-aea4-af64b83d78ab}")
)

type iFileSaveDialog struct {
	vtbl *iFileSaveDialogVtbl
}

type iFileSaveDialogVtbl struct {
	iFileDialogVtbl

	SetSaveAsItem          uintptr
	SetProperties          uintptr
	SetCollectedProperties uintptr
	GetProperties          uintptr
	ApplyProperties        uintptr
}

func newIFileSaveDialog() (*iFileSaveDialog, error) {
	if unknown, err := ole.CreateInstance(saveFileDialogCLSID, saveFileDialogIID); err == nil {
		return (*iFileSaveDialog)(unsafe.Pointer(unknown)), nil
	} else {
		return nil, err
	}
}

func (fileSaveDialog *iFileSaveDialog) Show() error {
	return fileSaveDialog.vtbl.show(unsafe.Pointer(fileSaveDialog))
}

func (fileSaveDialog *iFileSaveDialog) ShowAndGet() (string, error) {
	if err := fileSaveDialog.Show(); err != nil {
		return "", err
	}
	return fileSaveDialog.GetResult()
}

func (fileSaveDialog *iFileSaveDialog) Close() error {
	return fileSaveDialog.vtbl.close(unsafe.Pointer(fileSaveDialog))
}

func (fileSaveDialog *iFileSaveDialog) SetTitle(title string) error {
	return fileSaveDialog.vtbl.setTitle(unsafe.Pointer(fileSaveDialog), title)
}

func (fileSaveDialog *iFileSaveDialog) GetResult() (string, error) {
	return fileSaveDialog.vtbl.getResultString(unsafe.Pointer(fileSaveDialog))
}

func (fileSaveDialog *iFileSaveDialog) Release() error {
	return fileSaveDialog.vtbl.release(unsafe.Pointer(fileSaveDialog))
}

func (fileSaveDialog *iFileSaveDialog) SetDefaultFolder(defaultFolderPath string) error {
	return fileSaveDialog.vtbl.setDefaultFolder(unsafe.Pointer(fileSaveDialog), defaultFolderPath)
}

func (fileSaveDialog *iFileSaveDialog) SetInitialFolder(defaultFolderPath string) error {
	return fileSaveDialog.vtbl.setFolder(unsafe.Pointer(fileSaveDialog), defaultFolderPath)
}

func (fileSaveDialog *iFileSaveDialog) SetFileFilter(filter []FileFilter) error {
	return fileSaveDialog.vtbl.setFileTypes(unsafe.Pointer(fileSaveDialog), filter)
}

func (fileSaveDialog *iFileSaveDialog) SetRole(role string) error {
	return fileSaveDialog.vtbl.setClientGuid(unsafe.Pointer(fileSaveDialog), util.StringToUUID(role))
}
