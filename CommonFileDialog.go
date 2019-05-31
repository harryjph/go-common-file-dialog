// Cross-platform.

// Common File Dialogs
package cfd

func Initialize() error {
	return comInitialize()
}

func UnInitialize() {
	comUnInitialize()
}

type DialogConfig struct {
	// The title of the dialog
	Title string
	// The role of the dialog. This is used to derive the dialog's GUID, which the
	// OS will use to differentiate it from dialogs that are intended for other purposes.
	// This means that, for example, a dialog with role "Import" will have a different
	// previous location that it will open to than a dialog with role "Save". Can be any
	// string.
	Role string
	// The default folder - the folder that is used the first time the user opens it
	// (after the first time their last used location is used).
	DefaultFolder string
	// The initial folder - the folder that the dialog always opens to if not empty.
	// If this is not empty, it will override the "default folder" behaviour and
	// the dialog will always open to this folder.
	InitialFolder string
	// The file filter, in the format TODO
	FileFilter string
}

type Dialog interface {
	Show() error
	ShowAndGet() (string, error)
	Close() error // TODO does this even work?
	SetTitle(title string) error
	SetRole(role string) error
	SetDefaultFolder(defaultFolder string) error
	SetInitialFolder(folder string) error
	SetFileFilter(fileFilter string) error
	GetResult() (string, error)
	Release() error
}
