package util

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

// example alias:  alias := "\n# Added by dagger\nalias ll='ls -la'\n"

func AgnosticConfigUpdater(alias string) string {
	shellPath := ShellDefiner()

	switch {
	case strings.Contains(shellPath, "bash"):
		BashConfigUpdater(alias)
	case strings.Contains(shellPath, "zsh"):
		ZshConfigUpdater(alias)
	default:
		fmt.Println("Unsupported shell:", shellPath)
	}

	return "Configuration updated"
}

func ZshConfigUpdater(alias string) {
	usr, err := user.Current()
	if err != nil {
		fmt.Println("Error getting current user home dir:", err)
		return
	}

	zshConfigPath := filepath.Join(usr.HomeDir, ".zshrc")

	// Check if the file exists
	if _, err := os.Stat(zshConfigPath); os.IsNotExist(err) {
		fmt.Println("The .zshrc file does not exist at path:", zshConfigPath)
		return
	}

	configFile, err := os.OpenFile(zshConfigPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening zsh config file:", err)
		return
	}
	defer configFile.Close()

	if _, err := configFile.WriteString(alias); err != nil {
		fmt.Println("Error writing to zsh config file:", err)
		return
	}
}

func BashConfigUpdater(alias string) {
	usr, err := user.Current()
	if err != nil {
		fmt.Println("Error getting current user home dir:", err)
		return
	}

	bashConfigPath := filepath.Join(usr.HomeDir, ".bashrc")

	// Check if the file exists
	if _, err := os.Stat(bashConfigPath); os.IsNotExist(err) {
		fmt.Println("The .bashrc file does not exist at path:", bashConfigPath)
		return
	}

	configFile, err := os.OpenFile(bashConfigPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening bash config file:", err)
		return
	}
	defer configFile.Close()

	if _, err := configFile.WriteString(alias); err != nil {
		fmt.Println("Error writing to bash config file:", err)
		return
	}
}
