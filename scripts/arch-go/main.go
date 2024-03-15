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
    yay Script v0.0.2
`

var pacotes = []string{
	"adobe-source-han-sans-cn-fonts",
	"adobe-source-han-sans-jp-fonts",
	"adobe-source-han-sans-kr-fonts",
	"adobe-source-han-sans-otc-fonts",
	"adobe-source-han-sans-tw-fonts",
	"alacritty",
	"bitwarden",
	"pulseaudio",
	"pulseaudio-alsa",
	"alsa-utils",
	"lightdm",
	"lightdm-gtk-greeter",
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
	"openssh",
	"ranger",
	"rustup",
	"tmux",
	"ttf-hack-nerd",
	"ttf-inconsolata",
	"ttf-jetbrains-mono",
	"ttf-jetbrains-mono-nerd",
	"ttf-joypixels",
	"ttf-terminus-nerd",
	"ttf-unifont",
	"vim",
	"w3m",
	"wget",
	"zip",
	"zsh",
}

func main() {
	fmt.Printf("%s", logo)

	utils.InstallYay()
	fmt.Println("YAY INSTALLED")
	time.Sleep(1 * time.Second)
	utils.InstallLunarVim()

	for _, pacote := range pacotes {
		if !utils.IsInstalled(pacote) {
			utils.InstallPackage(pacote)
		} else {
			fmt.Printf("%s is already installed. ✅\n", pacote)
			time.Sleep(300 * time.Millisecond)
		}
	}

}
