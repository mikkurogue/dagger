package util

import (
	"fmt"
	"os"
	"runtime"
)

const (
	default_zsh_path  = "/usr/bin/zsh"
	default_bash_path = "/usr/bin/bash"
)

func ShellDefiner() string {

	var shell_name_dir string
	if runtime.GOOS == "windows" {
		shell_name_dir = os.Getenv("ComSpec")
		if shell_name_dir == "" {
			shell_name_dir = os.Getenv("SHELL")
		}
	} else {
		shell_name_dir = os.Getenv("SHELL")
	}

	if shell_name_dir == "" {
		fmt.Println("Could not determine the shell")
	}

	return string(shell_name_dir)
}