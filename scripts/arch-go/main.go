package main

import (
	"arch/utils"
	"fmt"
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
		if !utils.IsInstalled(pacote) {
			utils.InstallPackage(pacote)
		} else {
			fmt.Printf("%s is already installed. ✅\n", pacote)
			time.Sleep(1 * time.Second)
		}
	}
}
