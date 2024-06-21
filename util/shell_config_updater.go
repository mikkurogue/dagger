package util

import (
	"os"
	"os/user"
	"path/filepath"

	"github.com/fatih/color"
)

// example alias:  alias := "\n# Added by dagger\nalias ll='ls -la'\n"

func AgnosticConfigUpdater(alias string) {

		var shell_path = ShellDefiner()

		if shell_path == default_bash_path {
			BashConfigUpdater(alias)
		}

		if shell_path == default_zsh_path {
			ZshConfigUpdater(alias)
		}
	
}


func ZshConfigUpdater(alias string) {
	usr, err := user.Current()
	if err != nil {
		color.Red("Error getting current user home dir")
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

func BashConfigUpdater(alias string) {
	usr, err := user.Current()
	if err != nil {
		color.Red("Error getting current user home dir")
		return
	}

	bash_config_path := filepath.Join(usr.HomeDir, ".bashrc")

	config_file, err := os.OpenFile(bash_config_path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		color.Red("Error opening bash config file\n")
		return
	}
	defer config_file.Close()

	if _, err := config_file.WriteString(alias); err != nil {
		color.Red("Error updating bash config file\n")
		return
	}
}
