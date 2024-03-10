#!/bin/bash
#
CRE=$(tput setaf 1)
CYE=$(tput setaf 3)
CGR=$(tput setaf 2)
CBL=$(tput setaf 4)
BLD=$(tput bold)
CNC=$(tput sgr0)


logo () {
    local text="${1:?}"
    echo -en "
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
    yay Dotfiles\n\n"
    printf ' %s [%s%s %s%s %s]%s\n\n' "${CRE}" "${CNC}" "${CYE}" "${text}" "${CNC}" "${CRE}" "${CNC}"
}

logo "dança da bundinha peladinha"

pacman -Sy

printf "Installing neofetch, exa, tmux and bitwarden\n\n"

sleep 3

pacman -S --noconfirm neofetch exa tmux bitwarden alacritty rustup

printf "Installing asian fonts\n\n"

sleep 3

pacman -S --noconfirm adobe-source-han-sans-cn-fonts adobe-source-han-sans-jp-fonts adobe-source-han-sans-kr-fonts adobe-source-han-sans-otc-fonts adobe-source-han-sans-tw-fonts

