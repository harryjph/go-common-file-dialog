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
	GetResult() (string, error)
	Release() error
}

func NewOpenFileDialog() (Dialog, error) {
	return newIFileOpenDialog()
}

func NewSaveFileDialog() (Dialog, error) {
	return newIFileSaveDialog()
}
