package cfd

import (
	"github.com/go-ole/go-ole"
	"syscall"
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
	IID_FileOpenDialog   = "{D57C7288-D4AD-4768-BE02-9D969532D960}"
)

func newIFileOpenDialog() *iFileOpenDialog {
	err := ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED|ole.COINIT_DISABLE_OLE1DDE)
	if err == nil {
		openFileDialogClsid := clsid(CLSID_FileOpenDialog)
		openFileDialogIid := clsid(IID_FileOpenDialog)
		unknown, err := ole.CreateInstance(openFileDialogClsid, openFileDialogIid)
		if err != nil {
			panic(err)
		}
		fileOpen := (*iFileOpenDialog)(unsafe.Pointer(unknown))
		return fileOpen
	} else {
		panic(err)
	}
}

func clsid(str string) *ole.GUID { // TODO remove
	a, err := ole.CLSIDFromString(str)
	if err != nil {
		panic(err)
	}
	return a
}

func (fileOpenDialog *iFileOpenDialog) Show() uintptr {
	ret, _, _ := syscall.Syscall(fileOpenDialog.vtbl.Show,
		1,
		uintptr(unsafe.Pointer(fileOpenDialog)),
		0,
		0)
	return ret
}
