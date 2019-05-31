package cfd

type iUnknownVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
}

type iModalWindowVtbl struct {
	iUnknownVtbl
	Show uintptr // func (hwndOwner HWND) HRESULT
}

type iFileDialogVtbl struct {
	iModalWindowVtbl
	SetFileTypes        uintptr
	SetFileTypeIndex    uintptr
	GetFileTypeIndex    uintptr
	Advise              uintptr
	Unadvise            uintptr
	SetOptions          uintptr // func (fos FILEOPENDIALOGOPTIONS) HRESULT
	GetOptions          uintptr // func (pfos *FILEOPENDIALOGOPTIONS) HRESULT
	SetDefaultFolder    uintptr // func (psi *IShellItem) HRESULT
	SetFolder           uintptr
	GetFolder           uintptr
	GetCurrentSelection uintptr
	SetFileName         uintptr
	GetFileName         uintptr
	SetTitle            uintptr
	SetOkButtonLabel    uintptr
	SetFileNameLabel    uintptr
	GetResult           uintptr // func (ppsi **IShellItem) HRESULT
	AddPlace            uintptr
	SetDefaultExtension uintptr
	Close               uintptr // func (hr HRESULT) HRESULT
	SetClientGuid       uintptr
	ClearClientData     uintptr
	SetFilter           uintptr
}
