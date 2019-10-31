// +build windows

package cfd

import (
	"github.com/go-ole/go-ole"
	"sync"
)

var initialized = false
var initLock sync.Mutex

func comInitialize() {
	initLock.Lock()
	defer initLock.Unlock()

	if !initialized {
		ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED|ole.COINIT_DISABLE_OLE1DDE)
		initialized = true
	}
}

func comUnInitialize() {
	initLock.Lock()
	defer initLock.Unlock()

	ole.CoUninitialize()
	initialized = false
}
