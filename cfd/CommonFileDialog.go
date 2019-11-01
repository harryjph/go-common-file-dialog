// Cross-platform.

// Common File Dialogs
package cfd

type Dialog interface { // TODO setDefaultExtension?
	Show() error
	ShowAndGet() (string, error)
	Close() error // TODO does this even work?
	SetTitle(title string) error
	SetRole(role string) error
	SetDefaultFolder(defaultFolder string) error
	SetInitialFolder(folder string) error
	SetFileFilter(fileFilter []FileFilter) error
	GetResult() (string, error)
	Release() error
}

type OpenMultipleDialog interface {
	Dialog
	ShowAndGetAll() ([]string, error)
	GetResults() ([]string, error)
}
