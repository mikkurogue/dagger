package installer_windows

import (
	"dagger/util"
	"os"
	"os/exec"

	"github.com/charmbracelet/huh/spinner"
	"github.com/fatih/color"
)

func Tools(cli_tools []string) {
	for _, tool := range cli_tools {
		switch tool {
		case "lazygit":
			_ = spinner.New().Title("Installing lazygit..").Action(func() {
				_, err := exec.Command("winget", "install", "-e", "--id", "JesseDuffield.lazygit").Output()
				if err != nil {
					color.Red("Error installinglazygit\n")
					os.Exit(1)
				}
			}).Run()
		case "zoxide":
			_ = spinner.New().Title("Installing lazygit..").Action(func() {
				_, err := exec.Command("winget", "install", "ajeetdsouza.zoxide").Output()
				if err != nil {
					color.Red("Error installing zoxide\n")
					os.Exit(1)
				}

				util.AgnosticConfigUpdater("Invoke-Expression (& { (zoxide init powershell | Out-String) })")
			}).Run()
		case "skip":
			continue
		}
	}
}
