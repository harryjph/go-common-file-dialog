// Cross-platform.

// Common File Dialogs
package cfd

type Dialog interface {
	Show() error
	ShowAndGetResult() (string, error)
	Close() error // TODO does this even work?
	SetTitle(title string) error
	SetRole(role string) error
	SetDefaultFolder(defaultFolder string) error
	SetInitialFolder(folder string) error
	GetResult() (string, error)
	// For Select Folder Dialog, sets folder name.
	SetFileName(fileName string) error
	Release() error
}

type FileDialog interface {
	Dialog
	SetFileFilters(fileFilter []FileFilter) error
	SetSelectedFileFilterIndex(index uint) error
	SetDefaultExtension(defaultExtension string) error
}

type OpenFileDialog interface {
	FileDialog
}

type OpenMultipleFilesDialog interface {
	FileDialog
	ShowAndGetResults() ([]string, error)
	GetResults() ([]string, error)
}

type SelectFolderDialog interface {
	Dialog
}

type SaveFileDialog interface { // TODO Properties
	FileDialog
}
