package cfd

import (
	"github.com/go-ole/go-ole"
	"syscall"
	"unsafe"
)

func (vtbl *iShellItemVtbl) getDisplayName(objPtr unsafe.Pointer) (string, error) {
	const SIGDN_FILESYSPATH = 0x80058000
	var ptr *uint16
	ret, _, _ := syscall.Syscall(vtbl.GetDisplayName,
		2,
		uintptr(objPtr),
		SIGDN_FILESYSPATH,
		uintptr(unsafe.Pointer(&ptr)))
	if ret != 0 {
		return "", ole.NewError(ret)
	}
	return ole.LpOleStrToString(ptr), nil
}

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

func (vtbl *iFileDialogVtbl) getResult(objPtr unsafe.Pointer) (*iShellItem, error) {
	var shellItem iShellItem
	ret, _, _ := syscall.Syscall(vtbl.GetResult,
		1,
		uintptr(objPtr),
		uintptr(unsafe.Pointer(&shellItem)),
		0)
	if ret != 0 {
		return nil, ole.NewError(ret)
	}
	return &shellItem, nil
}
