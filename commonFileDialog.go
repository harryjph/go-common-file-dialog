// +build windows

// Common File Dialog for Windows
package cfd

type Dialog interface {
	Show() error
	Close() error
}

type OpenFileDialog interface {
	Dialog
	GetResult() string // TODO more methods
}

func NewOpenFileDialog() OpenFileDialog {
	return nil // TODO
}
