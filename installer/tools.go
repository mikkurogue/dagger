package installer

import (
	"dagger/util"
	"os"
	"os/exec"

	"github.com/charmbracelet/huh/spinner"
	"github.com/fatih/color"
)

func Tools(cli_tools []string, current_os string, curr_step int) {
	for _, tool := range cli_tools {
		switch tool {
		case "eza":
			_ = spinner.New().Title("Installing Eza...").Action(func() {
				curr_step += 1
				if current_os == "windows" {
					color.Red("can not install eza on this operating system\n")
					return
				}
				_, err := exec.Command("brew", "install", "eza").Output()
				if err != nil {
					color.Red("Error installing eza\n")
					os.Exit(1)
				}

				util.AgnosticConfigUpdater("\n# Added by dagger\nalias ls='eza --color=always --long --git --no-filesize --no-time --no-user --no-permissions --tree --level=2'")
			}).Run()
		case "fzf":
			_ = spinner.New().Title("Installing fzf...").Action(func() {
				curr_step += 1
				if current_os == "windows" {
					color.Red("can not install fzf on this operating system\n")
					return
				}
				_, err := exec.Command("brew", "install", "fzf").Output()
				if err != nil {
					color.Red("Error installing fzf\n")
					os.Exit(1)
				}
			}).Run()
		case "bat":
			_ = spinner.New().Title("Installing bat...").Action(func() {
				curr_step += 1
				if current_os == "windows" {
					color.Red("can not install bat on this operating system\n")
					return
				}
				_, err := exec.Command("brew", "install", "bat").Output()
				if err != nil {
					color.Red("Error installing bat\n")
					os.Exit(1)
				}
			}).Run()
		case "ripgrep":
			_ = spinner.New().Title("Installing Ripgrep...").Action(func() {
				curr_step += 1
				if current_os == "windows" {
					color.Red("can not install ripgrep on this operating system\n")
					return
				}
				_, err := exec.Command("brew", "install", "ripgrep").Output()
				if err != nil {
					color.Red("Error installing ripgrep\n")
					os.Exit(1)
				}
			}).Run()
		case "thefuck":
			_ = spinner.New().Title("Installing thefuck...").Action(func() {
				curr_step += 1
				if current_os != "darwin" {
					color.Red("can not install thefuck on this operating system\n")
					return
				}

				_, err := exec.Command("brew", "install", "thefuck").Output()
				if err != nil {
					color.Red("Error installing thefuck\n")
					os.Exit(1)
				}

				util.AgnosticConfigUpdater("\n# Added by dagger\neval $(thefuck --alias)")
			}).Run()
		case "lazygit":
			_ = spinner.New().Title("Installing lazygit..").Action(func() {
				curr_step += 1
				if current_os == "windows" {
					color.Red("can not install lazygit on this operating system using brew\n")
					return
				}
				_, err := exec.Command("brew", "install", "lazygit").Output()
				if err != nil {
					color.Red("Error installing oh my lazygit\n")
					os.Exit(1)
				}
			}).Run()
		case "skip":
			continue
		}
	}
}

func contains(s []string, str string) bool {
	for _, value := range s {
		if value == str {
			return true
		}
	}
	return false
}
