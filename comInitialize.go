package cfd

import "github.com/go-ole/go-ole"

func comInitialize() error {
	return ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED|ole.COINIT_DISABLE_OLE1DDE)
}
