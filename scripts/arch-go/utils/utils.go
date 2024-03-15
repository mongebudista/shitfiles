package utils

import (
	"fmt"
	"os"
	"os/exec"
  "time"
)

func IsInstalled(packageName string) bool {
	cmd := exec.Command("pacman", "-Q", packageName)
	err := cmd.Run()

	if err != nil {
		return false
	}
	return true
}

func InstallPackage(packageName string) {
	cmd := exec.Command("sudo", "pacman", "-S", "--noconfirm", packageName)
	err := cmd.Run()
	fmt.Printf("Installing %s... 📝\n", packageName)
	time.Sleep(1 * time.Second)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err != nil {
		fmt.Printf("Error while try install %s: %v", packageName, err)
		os.Exit(1)
	}
}
