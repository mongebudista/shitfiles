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
	"adobe-source-han-sans-cn-fonts",
	"adobe-source-han-sans-jp-fonts",
	"adobe-source-han-sans-kr-fonts",
	"adobe-source-han-sans-otc-fonts",
	"adobe-source-han-sans-tw-fonts",
	"alacritty",
	"base-devel",
  "curl",
	"exa",
  "feh",
	"firefox",
  "geany",
	"git",
	"go",
  "libwebp",
	"neofetch",
	"neovim",
  "pacman-contrib",
  "polybar",
  "picom",
  "ranger",
	"rustup",
	"tmux",
  "ttf-inconsolata",
  "ttf-jetbrains-mono", 
  "ttf-jetbrains-mono-nerd",
	"ttf-joypixels",
  "ttf-terminus-nerd",
  "vim",
  "w3m",
  "wget",
  "zip",
  "zsh",
}


func main() {
  fmt.Printf("%s", logo)

	for _, pacote := range pacotes {
		if !is_installed(pacote) {
			install_package(pacote)
		} else {
			fmt.Printf("%s is already installed. ✅\n", pacote)
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
	fmt.Printf("Installing %s... 📝\n", packageName)
  time.Sleep(1 * time.Second)

  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr

	if err != nil {
		fmt.Printf("Error while try install %s: %v", packageName, err)
		os.Exit(1)
	}

}
