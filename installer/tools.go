package installer

import (
	"dagger/util"
	"os"
	"os/exec"

	"github.com/charmbracelet/huh/spinner"
	"github.com/fatih/color"
)

func Tools(cli_tools []string, current_os string) {
	for _, tool := range cli_tools {
		switch tool {
		case "eza":
			_ = spinner.New().Title("Installing Eza...").Action(func() {
				_, err := exec.Command("brew", "install", "eza").Output()
				if err != nil {
					color.Red("Error installing eza\n")
					os.Exit(1)
				}
				util.AgnosticConfigUpdater("\n# Added by dagger\nalias ls=\"eza --color=always --long --git --no-filesize --no-time --no-user --no-permissions --tree --level=2\"")
			}).TitleStyle(util.TITLE_STYLE).Run()
		case "zoxide":
			_ = spinner.New().Title("Installing zoxide...").Action(func() {
				_, err := exec.Command("brew", "install", "zoxide").Output()
				if err != nil {
					color.Red("Error installing zoxide\n")
					os.Exit(1)
				}
				util.AgnosticConfigUpdater("\n# Added by dagger\nalias cd=\"z\"")
			}).TitleStyle(util.TITLE_STYLE).Run()
		case "fzf":
			_ = spinner.New().Title("Installing fzf...").Action(func() {
				_, err := exec.Command("brew", "install", "fzf").Output()
				if err != nil {
					color.Red("Error installing fzf\n")
					os.Exit(1)
				}
			}).TitleStyle(util.TITLE_STYLE).Run()
		case "bat":
			_ = spinner.New().Title("Installing bat...").Action(func() {
				_, err := exec.Command("brew", "install", "bat").Output()
				if err != nil {
					color.Red("Error installing bat\n")
					os.Exit(1)
				}
			}).TitleStyle(util.TITLE_STYLE).Run()
		case "ripgrep":
			_ = spinner.New().Title("Installing Ripgrep...").Action(func() {
				_, err := exec.Command("brew", "install", "ripgrep").Output()
				if err != nil {
					color.Red("Error installing ripgrep\n")
					os.Exit(1)
				}
			}).TitleStyle(util.TITLE_STYLE).Run()
		case "thefuck":
			_ = spinner.New().Title("Installing thefuck...").Action(func() {
				_, err := exec.Command("brew", "install", "thefuck").Output()
				if err != nil {
					color.Red("Error installing thefuck\n")
					os.Exit(1)
				}

				util.AgnosticConfigUpdater("\n# Added by dagger\n" + THE_FUCK_SETTING)
			}).TitleStyle(util.TITLE_STYLE).Run()
		case "lazygit":
			_ = spinner.New().Title("Installing lazygit..").Action(func() {
				_, err := exec.Command("brew", "install", "lazygit").Output()
				if err != nil {
					color.Red("Error installing lazygit\n")
					os.Exit(1)
				}
			}).TitleStyle(util.TITLE_STYLE).Run()
		case "skip":
			continue
		}
	}
}

const THE_FUCK_SETTING = "eval $(thefuck --alias)"

const EZA_SETTING = "alias ls='eza --color=always --long --git --no-filesize --no-time --no-user --no-permissions --tree --level=2'"
