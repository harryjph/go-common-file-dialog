package cfd

import (
	"github.com/go-ole/go-ole/oleutil"
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
	iidFileOpenDialog   = "{D57C7288-D4AD-4768-BE02-9D969532D960}" // TODO remove if unnedeed
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

func (fileOpenDialog *iFileOpenDialog) Close() error {
	return fileOpenDialog.vtbl.close(unsafe.Pointer(fileOpenDialog))
}

func (fileOpenDialog *iFileOpenDialog) GetResult() (string, error) {
	return fileOpenDialog.vtbl.getResultString(unsafe.Pointer(fileOpenDialog))
}

func (fileOpenDialog *iFileOpenDialog) Release() error {
	return fileOpenDialog.vtbl.release(unsafe.Pointer(fileOpenDialog))
}
