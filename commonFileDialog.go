// +build windows

// Common File Dialog for Windows
package cfd

import (
	"github.com/go-ole/go-ole"
	"unsafe"
)

const (
	CLSID_FileOpenDialog = "{DC1C5A9C-E88A-4dde-A5A1-60F82A20AEF7}"
	IID_FileOpenDialog   = "{D57C7288-D4AD-4768-BE02-9D969532D960}"
	CLSID_FileSaveDialog = "{C0B4E2F3-BA21-4773-8DBA-335EC946EB8B}"
)

func New() *iFileOpenDialog {
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

func clsid(str string) *ole.GUID {
	a, err := ole.CLSIDFromString(str)
	if err != nil {
		panic(err)
	}
	return a
}
