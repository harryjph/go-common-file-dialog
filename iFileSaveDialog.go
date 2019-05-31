// +build windows

package cfd

import (
	"github.com/go-ole/go-ole/oleutil"
	"unsafe"
)

const (
	clsidFilesavedialog = "{C0B4E2F3-BA21-4773-8DBA-335EC946EB8B}"
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
	if unknown, err := oleutil.CreateObject(clsidFilesavedialog); err == nil {
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

func (fileSaveDialog *iFileSaveDialog) SetFolder(defaultFolderPath string) error {
	return fileSaveDialog.vtbl.setFolder(unsafe.Pointer(fileSaveDialog), defaultFolderPath)
}

func (fileSaveDialog *iFileSaveDialog) SetFileFilter(defaultFolderPath string) error { // TODO
	return nil
}
