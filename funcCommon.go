package cfd

import (
	"github.com/go-ole/go-ole"
	"syscall"
	"unsafe"
)

func (vtbl *iModalWindowVtbl) show(objPtr unsafe.Pointer) error {
	ret, _, _ := syscall.Syscall(vtbl.Show,
		1,
		uintptr(objPtr),
		0,
		0)
	if ret != 0 {
		return ole.NewError(ret)
	}
	return nil
}

func (vtbl *iFileDialogVtbl) close(objPtr unsafe.Pointer) error {
	ret, _, _ := syscall.Syscall(vtbl.Close,
		1,
		uintptr(objPtr),
		0,
		0)
	if ret != 0 {
		return ole.NewError(ret)
	}
	return nil
}
