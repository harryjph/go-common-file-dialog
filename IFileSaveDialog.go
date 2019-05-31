package cfd

import (
	"github.com/go-ole/go-ole/oleutil"
	"unsafe"
)

const (
	clsidFilesavedialog = "{C0B4E2F3-BA21-4773-8DBA-335EC946EB8B}"
	iidFileSaveDialog   = "{84bccd23-5fde-4cdb-aea4-af64b83d78ab}" // TODO remove if unneeded
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

func (fileSaveDialog *iFileSaveDialog) Close() error {
	return fileSaveDialog.vtbl.close(unsafe.Pointer(fileSaveDialog))
}

func (fileSaveDialog *iFileSaveDialog) GetResult() (string, error) {
	shellItem, err := fileSaveDialog.vtbl.getResult(unsafe.Pointer(fileSaveDialog))
	if err != nil {
		return "", err
	}
	return shellItem.vtbl.getDisplayName(unsafe.Pointer(shellItem))
}
