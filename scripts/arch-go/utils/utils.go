package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
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

func InstallLunarVim() {
	fmt.Println("Installing LunarVim")
	cmd := exec.Command("sudo", "bash", "-c", `LV_BRANCH='release-1.3/neovim-0.9' bash <(curl -s https://raw.githubusercontent.com/LunarVim/LunarVim/release-1.3/neovim-0.9/utils/installer/install.sh)`)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		return
	}
}

func InstallYay() error {
	if _, err := exec.LookPath("git"); err != nil {
		return fmt.Errorf("git não encontrado no sistema")
	}

	if err := exec.Command("git", "clone", "https://aur.archlinux.org/yay.git").Run(); err != nil {
		return fmt.Errorf("falha ao clonar o repositório yay: %v", err)
	}

	if err := os.Chdir("yay"); err != nil {
		fmt.Println("KKKKKKKKK")
		return fmt.Errorf("falha ao acessar o diretório yay: %v", err)
	}

	makeCmd := exec.Command("makepkg", "-si")
	makeCmd.Stdout = os.Stdout
	makeCmd.Stderr = os.Stderr

	// Redireciona a entrada padrão para responder automaticamente com "Y" (sim) ao prompt
	makeCmd.Stdin = strings.NewReader("Y\n")

	if err := makeCmd.Run(); err != nil {
		return fmt.Errorf("falha ao compilar e instalar yay: %v", err)
	}

	if err := os.Chdir(".."); err != nil {
		return fmt.Errorf("Falha ao retornar para o diretorio anterior: %v", err)
	}

	if err := os.RemoveAll("yay"); err != nil {
		return fmt.Errorf("falha ao remover a pasta yay: %v", err)
	}

	return nil
}

func requireRoot() error {
	if os.Geteuid() != 0 {
		return fmt.Errorf("You're useless, execute as sudo")
	}
	return nil
}

func LoadConfigs() {
	homeDir, _ := os.UserHomeDir()

	fmt.Printf("Copying dotfiles...\n")
	time.Sleep(1 * time.Second)
	cmd := exec.Command("sudo", "cp", "-r", "./.config", homeDir)
	err := cmd.Run()

	if err != nil {
		return
	}
}
