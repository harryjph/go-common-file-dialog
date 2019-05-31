// +build windows

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
		1, // First argument is owner hWnd, we are just passing null
		uintptr(objPtr),
		0,
		0)
	return hresultToError(ret)
}

// Options are:
// FOS_OVERWRITEPROMPT	= 0x2,
// FOS_STRICTFILETYPES	= 0x4,
// FOS_NOCHANGEDIR	= 0x8,
// FOS_PICKFOLDERS	= 0x20,
// FOS_FORCEFILESYSTEM	= 0x40,
// FOS_ALLNONSTORAGEITEMS	= 0x80,
// FOS_NOVALIDATE	= 0x100,
// FOS_ALLOWMULTISELECT	= 0x200,
// FOS_PATHMUSTEXIST	= 0x800,
// FOS_FILEMUSTEXIST	= 0x1000,
// FOS_CREATEPROMPT	= 0x2000,
// FOS_SHAREAWARE	= 0x4000,
// FOS_NOREADONLYRETURN	= 0x8000,
// FOS_NOTESTFILECREATE	= 0x10000,
// FOS_HIDEMRUPLACES	= 0x20000,
// FOS_HIDEPINNEDPLACES	= 0x40000,
// FOS_NODEREFERENCELINKS	= 0x100000,
// FOS_OKBUTTONNEEDSINTERACTION	= 0x200000,
// FOS_DONTADDTORECENT	= 0x2000000,
// FOS_FORCESHOWHIDDEN	= 0x10000000,
// FOS_DEFAULTNOMINIMODE	= 0x20000000,
// FOS_FORCEPREVIEWPANEON	= 0x40000000,
// FOS_SUPPORTSTREAMABLEITEMS	= 0x80000000
func (vtbl *iFileDialogVtbl) setOptions(objPtr unsafe.Pointer, options uint32) error {
	ret, _, _ := syscall.Syscall(vtbl.SetOptions,
		1,
		uintptr(objPtr),
		uintptr(options),
		0)
	return hresultToError(ret)
}

func (vtbl *iFileDialogVtbl) getOptions(objPtr unsafe.Pointer) (uint32, error) {
	var options uint32
	ret, _, _ := syscall.Syscall(vtbl.GetOptions,
		1,
		uintptr(objPtr),
		uintptr(unsafe.Pointer(&options)),
		0)
	return options, hresultToError(ret)
}

func (vtbl *iFileDialogVtbl) addOption(objPtr unsafe.Pointer, option uint32) error {
	if options, err := vtbl.getOptions(objPtr); err == nil {
		return vtbl.setOptions(objPtr, options|option)
	} else {
		return err
	}
}

func (vtbl *iFileDialogVtbl) removeOption(objPtr unsafe.Pointer, option uint32) error {
	if options, err := vtbl.getOptions(objPtr); err == nil {
		return vtbl.setOptions(objPtr, options&^option)
	} else {
		return err
	}
}

func (vtbl *iFileDialogVtbl) setDefaultFolder(objPtr unsafe.Pointer, path string) error {
	shellItem, err := newIShellItem(path) // TODO do we need to defer release()
	if err != nil {
		return err
	}
	ret, _, _ := syscall.Syscall(vtbl.SetDefaultFolder,
		1,
		uintptr(objPtr),
		uintptr(unsafe.Pointer(shellItem)),
		0)
	return hresultToError(ret)
}

func (vtbl *iFileDialogVtbl) setTitle(objPtr unsafe.Pointer, title string) error {
	titlePtr := ole.SysAllocString(title) // TODO do we need to CoTaskMemFree?
	ret, _, _ := syscall.Syscall(vtbl.SetTitle,
		1,
		uintptr(objPtr),
		uintptr(unsafe.Pointer(titlePtr)),
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
