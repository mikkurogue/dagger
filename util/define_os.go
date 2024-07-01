package util

import "runtime"

func DefineOs(currentOs *string) *string {
	if runtime.GOOS == "windows" {
		*currentOs = "windows"
	}

	if runtime.GOOS == "darwin" {
		*currentOs = "darwin"
	}

	if runtime.GOOS == "linux" {
		*currentOs = "linux"
	}

	return currentOs
}