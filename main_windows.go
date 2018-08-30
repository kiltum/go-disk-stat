package main

import (
	win "golang.org/x/sys/windows"
	"unsafe"
)

func DiskFree(path string) (uint64, uint64, error) {

	c := win.MustLoadDLL("kernel32.dll").MustFindProc("GetDiskFreeSpaceExW")

	pointer, err := win.UTF16PtrFromString(path)

	if err != nil {
		return 0, 0, err
	}

	var freeBytes uint64
	var totalBytes uint64
	var availBytes uint64

	c.Call(
		uintptr(unsafe.Pointer(pointer)),
		uintptr(unsafe.Pointer(&freeBytes)),
		uintptr(unsafe.Pointer(&totalBytes)),
		uintptr(unsafe.Pointer(&availBytes)))

	if err != nil {
		return 0, 0, err
	}
	return freeBytes, totalBytes, nil
}
