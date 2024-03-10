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

clear

printf "Installing packages...\n\n"
sleep 3

# Array com os nomes dos pacotes
pacotes=(
  "neofetch"
  "bitwarden"
  "exa"
  "alacritty"
  "rustup"
  "tmux"
  "neovim"
  "git"
  "go"
  "adobe-source-han-sans-cn-fonts"
  "adobe-source-han-sans-jp-fonts"
  "adobe-source-han-sans-kr-fonts"
  "adobe-source-han-sans-otc-fonts"
  "adobe-source-han-sans-tw-fonts"
)

# Função para verificar se o pacote está instalado e instalá-lo se não estiver
instalar_pacote() {
    if ! command -v "$1" &> /dev/null; then
        echo "Instalando $1..."
        sudo pacman -S --noconfirm "$1"
    else
        echo "$1 já está instalado."
    fi
}

# Itera sobre o array de pacotes e verifica/instala cada um
for pacote in "${pacotes[@]}"; do
    instalar_pacote "$pacote"
done

clear

logo 'OK!'
sleep 1
