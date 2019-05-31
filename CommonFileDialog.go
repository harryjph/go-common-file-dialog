// Common File Dialogs
package cfd

func Initialize() error {
	return comInitialize()
}

func UnInitialize() {
	comUnInitialize()
}

type Dialog interface {
	Show() error
	ShowAndGet() (string, error)
	Close() error // TODO does this even work?
	SetTitle(title string) error
	SetDefaultFolder(defaultFolder string) error
	SetFolder(folder string) error
	SetFileFilter(fileFilter string) error
	GetResult() (string, error)
	Release() error
}
