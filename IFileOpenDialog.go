package cfd

import (
	"github.com/go-ole/go-ole"
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
	CLSID_FileOpenDialog = "{DC1C5A9C-E88A-4dde-A5A1-60F82A20AEF7}"
	IID_FileOpenDialog   = "{D57C7288-D4AD-4768-BE02-9D969532D960}" // TODO remove if unnedeed
)

func newIFileOpenDialog() (*iFileOpenDialog, error) {
	if err := ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED|ole.COINIT_DISABLE_OLE1DDE); err != nil {
		return nil, err
	}
	if unknown, err := oleutil.CreateObject(CLSID_FileOpenDialog); err == nil {
		return (*iFileOpenDialog)(unsafe.Pointer(unknown)), nil
	} else {
		return nil, err
	}
}

func (fileOpenDialog *iFileOpenDialog) Show() error {
	return fileOpenDialog.vtbl.show(unsafe.Pointer(fileOpenDialog))
}
