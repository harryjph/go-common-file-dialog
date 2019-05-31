package cfd

import (
	"github.com/go-ole/go-ole"
	"syscall"
	"unsafe"
)

func hresultToError(hr uintptr) error {
	if hr != 0 {
		return ole.NewError(hr)
	}
	return nil
}

func (vtbl *iUnknownVtbl) release(objPtr unsafe.Pointer) error {
	ret, _, _ := syscall.Syscall(vtbl.Release,
		0,
		uintptr(objPtr),
		0,
		0)
	return hresultToError(ret)
}

func (vtbl *iModalWindowVtbl) show(objPtr unsafe.Pointer) error {
	ret, _, _ := syscall.Syscall(vtbl.Show,
		1,
		uintptr(objPtr),
		0,
		0)
	return hresultToError(ret)
}

func (vtbl *iFileDialogVtbl) close(objPtr unsafe.Pointer) error {
	ret, _, _ := syscall.Syscall(vtbl.Close,
		1,
		uintptr(objPtr),
		0,
		0)
	return hresultToError(ret)
}

func (vtbl *iFileDialogVtbl) getResult(objPtr unsafe.Pointer) (*iShellItem, error) {
	var shellItem *iShellItem
	ret, _, _ := syscall.Syscall(vtbl.GetResult,
		1,
		uintptr(objPtr),
		uintptr(unsafe.Pointer(&shellItem)),
		0)
	return shellItem, hresultToError(ret)
}

func (vtbl *iFileDialogVtbl) getResultString(objPtr unsafe.Pointer) (string, error) {
	shellItem, err := vtbl.getResult(objPtr)
	if err != nil {
		return "", err
	}
	defer shellItem.vtbl.release(unsafe.Pointer(shellItem))
	return shellItem.vtbl.getDisplayName(unsafe.Pointer(shellItem))
}
