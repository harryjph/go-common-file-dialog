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
}

type OpenFileDialog interface {
	Dialog
	GetResult() string // TODO more methods
}

type SaveFileDialog interface {
	Dialog
	// TODO more methods
}

func NewSaveFileDialog() OpenFileDialog {
	return nil // TODO
}

func NewOpenFileDialog() OpenFileDialog {
	return nil // TODO
}
