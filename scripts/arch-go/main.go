package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

var logo = `
    ⣿⣿⣿⣿⣿⣿⣿⣿⠟⠁⠀⠀⠀⠀⣤⣤⡄⠀⠀⠀⠀⠙⣿⣿⣿⣿⣿⣿⣿⣿ 
    ⣿⣿⣿⣿⣿⣿⡿⠋⠀⠀⠔⠓⠄⣀⣿⣿⣧⠄⣀⣍⡂⠀⡸⣿⣿⣿⣿⣿⣿⣿ 
    ⣿⣿⣿⣿⣿⠟⠀⠠⣡⣶⠿⠿⠿⣎⣻⣿⣇⠾⠟⠿⢿⣧⢱⡸⣿⣿⣿⣿⣿⣿ 
    ⣿⣿⣿⠿⠃⠀⢀⣴⠟⠀⢀⣀⡀⠈⢿⣿⠁⠀⣀⡀⠀⠹⣷⣷⣶⣶⣬⣝⢿⣿ 
    ⣿⡿⢣⣾⣿⣄⢸⣿⡀⠀⠘⠛⠃⠀⣸⣿⡀⠀⠛⠃⠀⢠⣿⣿⣿⣿⣿⣿⣦⢻ 
    ⣿⢣⣿⣿⣿⣿⣿⣿⣷⣤⣤⣤⣤⣼⣿⣿⣷⣤⣤⣤⣴⣿⣿⣿⣿⣿⣿⣿⣿⡝ 
    ⣿⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢁ 
    ⣿⣷⣬⣭⣥⣶⣬⣙⡛⠻⢿⣿⣿⣿⣿⣿⣿⣿⠋⠀⠈⠻⣿⣿⣿⣿⣿⣿⡟⣸ 
    ⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣶⣶⣬⣍⣙⠻⠿⠿⠀⠀⠀⠀⢻⣿⣿⣿⠿⢋⣼⣿ 
    ⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣶⣄⡀⠀⣀⣬⣭⣤⣶⣶⣿⣿⣿
    yay Script v0.0.1
`

var pacotes = []string{
	"base-devel",
	"bitwarden",
	"firefox",
	"exa",
  "zsh",
	"neofetch",
	"alacritty",
	"rustup",
	"tmux",
	"git",
	"go",
	"neovim",
  "vim",
  "jre8",
  "ranger",
  "w3m",
  "wget",
  "curl",
  "zip",
  "fortune",
	"adobe-source-han-sans-cn-fonts",
	"adobe-source-han-sans-jp-fonts",
	"adobe-source-han-sans-kr-fonts",
	"adobe-source-han-sans-otc-fonts",
	"adobe-source-han-sans-tw-fonts",
}


func main() {
  fmt.Printf("%s", logo)

	for _, pacote := range pacotes {
		if !is_installed(pacote) {
			install_package(pacote)
		} else {
			fmt.Printf("%s já está instalado ✅\n", pacote)
			time.Sleep(1 * time.Second)
		}
	}
}

func is_installed(packageName string) bool {
	cmd := exec.Command("pacman", "-Q", packageName)
  err := cmd.Run()

	if err != nil {
		return false
	}
	return true
}

func install_package(packageName string) {
	cmd := exec.Command("sudo", "pacman", "-S", "--noconfirm", packageName)
  err := cmd.Run()
	fmt.Printf("Instalando %s...", packageName)
  fmt.Printf("%s", logo)

	if err != nil {
		fmt.Printf("Erro ao instalar %s: %v", packageName, err)
		os.Exit(1)
	}

}
