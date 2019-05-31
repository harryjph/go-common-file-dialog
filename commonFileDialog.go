// +build windows

// Common File Dialog for Windows
package cfd

import "log"

func init() {
	if err := comInitialize(); err != nil {
		log.Fatal(err) // TODO don't fatal
	}
}

type Dialog interface {
	Show() error
	Close() error
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
