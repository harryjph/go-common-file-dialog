package cfd

type iUnknownVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
}

type iShellItem struct { // TODO move this and its logic into separate file?
	vtbl *iShellItemVtbl
}

type iShellItemVtbl struct {
	iUnknownVtbl
	BindToHandler  uintptr
	GetParent      uintptr
	GetDisplayName uintptr // func (sigdnName SIGDN, ppszName *LPWSTR) HRESULT
	GetAttributes  uintptr
	Compare        uintptr
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
	SetOptions          uintptr
	GetOptions          uintptr
	SetDefaultFolder    uintptr
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
