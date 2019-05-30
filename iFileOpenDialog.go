package cfd

import (
	"syscall"
	"unsafe"
)

type iFileOpenDialog struct {
	vtbl *iFileOpenDialogVtbl
}

type iFileOpenDialogVtbl struct {
	// IUnknown
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	// IModalWindow
	Show uintptr

	// IFileDialog
	SetFileTypes        uintptr
	SetFileTypeIndex    uintptr
	GetFileTypeIndex    uintptr
	Advise              uintptr
	Unadvise            uintptr
	SetOptions          uintptr
	GetOptions          uintptr
	SetDefaultFolder    uintptr
	SetFolder           uintptr
	GetFolder           uintptr
	GetCurrentSelection uintptr
	SetFileName         uintptr
	GetFileName         uintptr
	SetTitle            uintptr
	SetOkButtonLabel    uintptr
	SetFileNameLabel    uintptr
	GetResult           uintptr
	AddPlace            uintptr
	SetDefaultExtension uintptr
	Close               uintptr
	SetClientGuid       uintptr
	ClearClientData     uintptr
	SetFilter           uintptr

	// IFileOpenDialog
	GetResults       uintptr
	GetSelectedItems uintptr
}

func (fileOpenDialog *iFileOpenDialog) Show() uintptr {
	ret, _, _ := syscall.Syscall(fileOpenDialog.vtbl.Show,
		2,
		uintptr(unsafe.Pointer(fileOpenDialog)),
		0,
		0)
	return ret
}
