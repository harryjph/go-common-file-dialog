// +build windows

package cfd

import (
	"fmt"
	"github.com/go-ole/go-ole"
	"github.com/harry1453/go-common-file-dialog/util"
	"syscall"
	"unsafe"
)

var (
	fileOpenDialogCLSID = ole.NewGUID("{DC1C5A9C-E88A-4dde-A5A1-60F82A20AEF7}")
	fileOpenDialogIID   = ole.NewGUID("{d57c7288-d4ad-4768-be02-9d969532d960}")
)

type iFileOpenDialog struct {
	vtbl *iFileOpenDialogVtbl
}

type iFileOpenDialogVtbl struct {
	iFileDialogVtbl

	GetResults       uintptr // func (ppenum **IShellItemArray) HRESULT
	GetSelectedItems uintptr
}

func newIFileOpenDialog() (*iFileOpenDialog, error) {
	if unknown, err := ole.CreateInstance(fileOpenDialogCLSID, fileOpenDialogIID); err == nil {
		return (*iFileOpenDialog)(unsafe.Pointer(unknown)), nil
	} else {
		return nil, err
	}
}

func (fileOpenDialog *iFileOpenDialog) Show() error {
	return fileOpenDialog.vtbl.show(unsafe.Pointer(fileOpenDialog))
}

func (fileOpenDialog *iFileOpenDialog) ShowAndGet() (string, error) {
	if err := fileOpenDialog.Show(); err != nil {
		return "", err
	}
	return fileOpenDialog.GetResult()
}

func (fileOpenDialog *iFileOpenDialog) ShowAndGetAll() ([]string, error) {
	if err := fileOpenDialog.Show(); err != nil {
		return nil, err
	}
	return fileOpenDialog.GetResults()
}

func (fileOpenDialog *iFileOpenDialog) Close() error {
	return fileOpenDialog.vtbl.close(unsafe.Pointer(fileOpenDialog))
}

func (fileOpenDialog *iFileOpenDialog) SetTitle(title string) error {
	return fileOpenDialog.vtbl.setTitle(unsafe.Pointer(fileOpenDialog), title)
}

func (fileOpenDialog *iFileOpenDialog) GetResult() (string, error) {
	return fileOpenDialog.vtbl.getResultString(unsafe.Pointer(fileOpenDialog))
}

func (fileOpenDialog *iFileOpenDialog) Release() error {
	return fileOpenDialog.vtbl.release(unsafe.Pointer(fileOpenDialog))
}

func (fileOpenDialog *iFileOpenDialog) SetDefaultFolder(defaultFolderPath string) error {
	return fileOpenDialog.vtbl.setDefaultFolder(unsafe.Pointer(fileOpenDialog), defaultFolderPath)
}

func (fileOpenDialog *iFileOpenDialog) SetInitialFolder(defaultFolderPath string) error {
	return fileOpenDialog.vtbl.setFolder(unsafe.Pointer(fileOpenDialog), defaultFolderPath)
}

func (fileOpenDialog *iFileOpenDialog) SetFileFilter(filter []FileFilter) error {
	return fileOpenDialog.vtbl.setFileTypes(unsafe.Pointer(fileOpenDialog), filter)
}

func (fileOpenDialog *iFileOpenDialog) SetRole(role string) error {
	return fileOpenDialog.vtbl.setClientGuid(unsafe.Pointer(fileOpenDialog), util.StringToUUID(role))
}

// This should only be callable when the user asks for a multi select because
// otherwise they will be given the Dialog interface which does not expose this function.
func (fileOpenDialog *iFileOpenDialog) GetResults() ([]string, error) {
	return fileOpenDialog.vtbl.getResultsStrings(unsafe.Pointer(fileOpenDialog))
}

func (fileOpenDialog *iFileOpenDialog) setPickFolders(pickFolders bool) error {
	const FosPickfolders = 0x20
	if pickFolders {
		return fileOpenDialog.vtbl.addOption(unsafe.Pointer(fileOpenDialog), FosPickfolders)
	} else {
		return fileOpenDialog.vtbl.removeOption(unsafe.Pointer(fileOpenDialog), FosPickfolders)
	}
}

func (fileOpenDialog *iFileOpenDialog) setIsMultiselect(isMultiSelect bool) error {
	const FosAllowMultiselect = 0x200
	if isMultiSelect {
		return fileOpenDialog.vtbl.addOption(unsafe.Pointer(fileOpenDialog), FosAllowMultiselect)
	} else {
		return fileOpenDialog.vtbl.removeOption(unsafe.Pointer(fileOpenDialog), FosAllowMultiselect)
	}
}

func (vtbl *iFileOpenDialogVtbl) getResults(objPtr unsafe.Pointer) (*iShellItemArray, error) {
	var shellItemArray *iShellItemArray
	ret, _, _ := syscall.Syscall(vtbl.GetResults,
		1,
		uintptr(objPtr),
		uintptr(unsafe.Pointer(&shellItemArray)),
		0)
	return shellItemArray, hresultToError(ret)
}

func (vtbl *iFileOpenDialogVtbl) getResultsStrings(objPtr unsafe.Pointer) ([]string, error) {
	shellItemArray, err := vtbl.getResults(objPtr)
	if err != nil {
		return nil, err
	}
	if shellItemArray == nil {
		return nil, fmt.Errorf("ShellItemArray was nil")
	}
	defer shellItemArray.vtbl.release(unsafe.Pointer(shellItemArray))
	count, err := shellItemArray.vtbl.getCount(unsafe.Pointer(shellItemArray))
	if err != nil {
		return nil, err
	}
	var results []string
	for i := uintptr(0); i < count; i++ {
		newItem, err := shellItemArray.vtbl.getItemAt(unsafe.Pointer(shellItemArray), i)
		if err != nil {
			return nil, err
		}
		results = append(results, newItem)
	}
	return results, nil
}
