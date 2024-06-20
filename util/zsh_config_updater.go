package util

import (
	"os"
	"os/user"
	"path/filepath"

	"github.com/fatih/color"
)

func ZshConfigUpdater(alias string) {
	usr, err := user.Current()
	if err != nil {
		color.Red("Error getting current usern")
		return
	}

	zsh_config_path := filepath.Join(usr.HomeDir, ".zshrc")

	config_file, err := os.OpenFile(zsh_config_path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		color.Red("Error opening zsh config file\n")
		return
	}
	defer config_file.Close()

	if _, err := config_file.WriteString(alias); err != nil {
		color.Red("Error updating zsh config file\n")
		return
	}

}
