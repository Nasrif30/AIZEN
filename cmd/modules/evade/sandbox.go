package evade

import (
	"runtime"
	"syscall"
	"time"
	"unsafe"
)

var (
	kernel32            = syscall.NewLazyDLL("kernel32.dll")
	GlobalMemoryStatusEx = kernel32.NewProc("GlobalMemoryStatusEx")
	GetTickCount64      = kernel32.NewProc("GetTickCount64")
)

type MEMORYSTATUSEX struct {
	Length               uint32
	MemoryLoad           uint32
	TotalPhys            uint64
	AvailPhys            uint64
	TotalPageFile        uint64
	AvailPageFile        uint64
	TotalVirtual         uint64
	AvailVirtual         uint64
	AvailExtendedVirtual uint64
}

func IsSandboxed() bool {
	if runtime.NumCPU() < 2 {
		return true
	}
	if getTotalRAM() < 4*1024*1024*1024 {
		return true
	}
	if getUptime() < 30*60 {
		return true
	}
	if hasVMArtifacts() {
		return true
	}
	return false
}

func getTotalRAM() uint64 {
	var state MEMORYSTATUSEX
	state.Length = uint32(unsafe.Sizeof(state))
	ret, _, _ := GlobalMemoryStatusEx.Call(uintptr(unsafe.Pointer(&state)))
	if ret == 0 {
		return 0
	}
	return state.TotalPhys
}

func getUptime() int64 {
	ret, _, _ := GetTickCount64.Call()
	return int64(ret / 1000)
}

func hasVMArtifacts() bool {
	return false
}