package installer

import (
	"os"
	"os/exec"

	"github.com/charmbracelet/huh/spinner"
	"github.com/fatih/color"
)

func Tools(cli_tools []string, current_os string, curr_step int) {

	containsOhMyZsh := contains(cli_tools, "oh-my-zsh")

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

				// check if there is omz in the selection, then skip setting this for now
				// and in omz setting check if eza was selected and append to .zshrc
				if !containsOhMyZsh {
					_, setupErr := exec.Command("/bin/zsh", "-c", "alias ls='eza --color=always --long --git --no-filesize --no-time --no-user --no-permission --tree --level=2'").Output()
					if setupErr != nil {
						color.Red("Can not set ls alias\n")
					}
				}
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
		case "oh-my-zsh":
			_ = spinner.New().Title("Installing Oh my zsh...").Action(func() {
				curr_step += 1
				if current_os == "windows" {
					color.Red("can not install oh my zsh on this operating system\n")
					return
				}
				_, err := exec.Command("brew", "install", "oh-my-zsh").Output()
				if err != nil {
					color.Red("Error installing oh my zsh\n")
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

				_, setupErr := exec.Command("/bin/zsh", "-c", "eval $(thefuck --alias)").Output()
				if setupErr != nil {
					color.Red("Can not set thefuck alias\n")
				}
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
