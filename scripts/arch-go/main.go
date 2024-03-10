package main

import (
	"fmt"
	"os"
	"os/exec"
  "time"
)

var pacotes = []string{
  "base-devel",
	"bitwarden",
	"exa",
	"neofetch",
	"kitty",
	"rustup",
	"tmux",
  "git",
  "go",
  "neovim",
  "adobe-source-han-sans-cn-fonts",
  "adobe-source-han-sans-jp-fonts",
  "adobe-source-han-sans-kr-fonts",
  "adobe-source-han-sans-otc-fonts",
  "adobe-source-han-sans-tw-fonts",
  "firefox",
}

func main() {
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

  if err := cmd.Run(); err != nil {
    return false
  }
  return true
}

func install_package(packageName string) {
  fmt.Printf("Instalando %s...", packageName)
  cmd := exec.Command("sudo", "pacman", "-S", "--noconfirm", packageName)
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  
  if err :=  cmd.Run(); err != nil {
    fmt.Printf("Erro ao instalar %s: %v", packageName, err)
    os.Exit(1)
  }
}
