package cfd

import (
	"github.com/go-ole/go-ole"
	"syscall"
	"unsafe"
)

type iShellItem struct {
	vtbl *iShellItemVtbl
}

type iShellItemVtbl struct {
	iUnknownVtbl
	BindToHandler  uintptr
	GetParent      uintptr
	GetDisplayName uintptr // func (sigdnName SIGDN, ppszName *LPWSTR) HRESULT
	GetAttributes  uintptr
	Compare        uintptr
}

func (vtbl *iShellItemVtbl) getDisplayName(objPtr unsafe.Pointer) (string, error) {
	var ptr *uint16
	ret, _, _ := syscall.Syscall(vtbl.GetDisplayName,
		2,
		uintptr(objPtr),
		0x80058000, // SIGDN_FILESYSPATH
		uintptr(unsafe.Pointer(&ptr)))
	if ret != 0 {
		return "", ole.NewError(ret)
	}
	return ole.LpOleStrToString(ptr), nil
}
