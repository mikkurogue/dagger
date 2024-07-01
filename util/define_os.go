package util

import "runtime"

func DefineOs(current_os *string) *string {
	if runtime.GOOS == "windows" {
		*current_os = "windows"
	}

	if runtime.GOOS == "darwin" {
		*current_os = "darwin"
	}

	if runtime.GOOS == "linux" {
		*current_os = "linux"
	}

	return current_os
}