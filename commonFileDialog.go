// +build windows

// Common File Dialog for Windows
package cfd

func Initialize() error {
	return comInitialize()
}

func UnInitialize() {
	comUnInitialize()
}

type Dialog interface { // TODO set title
	Show() error
	Close() error // TODO does this even work?
	SetDefaultFolder(defaultFolder string) error
	GetResult() (string, error)
	Release() error
}

type OpenDialog interface {
	Dialog
	SetPickFolders(bool) error
}

type SaveDialog interface {
	Dialog
}

func NewOpenFileDialog() (OpenDialog, error) {
	return newIFileOpenDialog()
}

func NewSaveFileDialog() (SaveDialog, error) {
	return newIFileSaveDialog()
}
