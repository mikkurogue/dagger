package util

import (
	"os/user"
	"path/filepath"
)

// For now assuming using PowerShell Core - add maybe other powershells later?
func GetPowershellProfilePath() (string, any) {
	usr, err := user.Current()
	if err != nil {
		return "Something went wrong getting the current user.", nil
	}

	// Assuming PowerShell Core (you might need to adjust for older versions)
	profilePath := filepath.Join(usr.HomeDir, "Documents", "PowerShell", "Microsoft.PowerShell_profile.ps1")
	return profilePath, "success"
}
