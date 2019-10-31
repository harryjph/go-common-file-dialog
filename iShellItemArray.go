package cfd

import (
	"syscall"
	"unsafe"
)

type iShellItemArray struct {
	vtbl *iShellItemArrayVtbl
}

type iShellItemArrayVtbl struct {
	iUnknownVtbl

	BindToHandler              uintptr
	GetPropertyStore           uintptr
	GetPropertyDescriptionList uintptr
	GetAttributes              uintptr
	GetCount                   uintptr // func (pdwNumItems *DWORD) HRESULT
	GetItemAt                  uintptr // func (dwIndex DWORD, ppsi **IShellItem) HRESULT
	EnumItems                  uintptr
}

func (vtbl *iShellItemArrayVtbl) getCount(objPtr unsafe.Pointer) (int32, error) {
	// TODO
	return 0, nil
}

func (vtbl *iShellItemArrayVtbl) getItemAt(objPtr unsafe.Pointer, index int32) (string, error) {
	var shellItem *iShellItem
	ret, _, _ := syscall.Syscall(vtbl.GetItemAt,
		2,
		uintptr(objPtr),
		uintptr(index),
		uintptr(unsafe.Pointer(&shellItem)))
	if err := hresultToError(ret); err != nil {
		return "", err
	}
	// TODO nil check on shellItem
	defer shellItem.vtbl.release(unsafe.Pointer(shellItem))
	return shellItem.vtbl.getDisplayName(unsafe.Pointer(shellItem)) // TODO do we really need to wrap every single call using a pointer?
}
