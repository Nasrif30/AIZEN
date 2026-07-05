package evade

import (
	"syscall"
	"unsafe"
)

var (
	ntdll                    = syscall.NewLazyDLL("ntdll.dll")
	NtCreateThreadExProc     = ntdll.NewProc("NtCreateThreadEx")
	NtAllocateVirtualMemoryProc = ntdll.NewProc("NtAllocateVirtualMemory")
	NtWriteVirtualMemoryProc = ntdll.NewProc("NtWriteVirtualMemory")
	NtOpenProcessProc        = ntdll.NewProc("NtOpenProcess")
)

func NtCreateThreadEx(process syscall.Handle, startAddress uintptr, parameter uintptr) (syscall.Handle, error) {
	var thread syscall.Handle
	status, _, _ := NtCreateThreadExProc.Call(
		uintptr(unsafe.Pointer(&thread)),
		0x1F0FFF,
		0,
		uintptr(process),
		startAddress,
		parameter,
		0,
		0,
		0,
		0,
		0,
	)
	if status != 0 {
		return 0, syscall.Errno(status)
	}
	return thread, nil
}

func VirtualAllocEx(process syscall.Handle, address uintptr, size int, allocationType, protect uint32) (uintptr, error) {
	var addr uintptr
	status, _, _ := NtAllocateVirtualMemoryProc.Call(
		uintptr(process),
		uintptr(unsafe.Pointer(&addr)),
		0,
		uintptr(unsafe.Pointer(&size)),
		uintptr(allocationType),
		uintptr(protect),
	)
	if status != 0 {
		return 0, syscall.Errno(status)
	}
	return addr, nil
}

func WriteProcessMemory(process syscall.Handle, address uintptr, data []byte) error {
	var written uintptr
	status, _, _ := NtWriteVirtualMemoryProc.Call(
		uintptr(process),
		address,
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(len(data)),
		uintptr(unsafe.Pointer(&written)),
	)
	if status != 0 {
		return syscall.Errno(status)
	}
	return nil
}

func OpenProcess(desiredAccess uint32, inheritHandle bool, processID uint32) (syscall.Handle, error) {
	var handle syscall.Handle
	inherit := uintptr(0)
	if inheritHandle {
		inherit = 1
	}
	status, _, _ := NtOpenProcessProc.Call(
		uintptr(unsafe.Pointer(&handle)),
		uintptr(desiredAccess),
		inherit,
		uintptr(processID),
	)
	if status != 0 {
		return 0, syscall.Errno(status)
	}
	return handle, nil
}