package installer

import (
	"dagger/util"
	"log"
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
					color.Red("Error installing eza\n" + EZA_SETTING)
					os.Exit(1)
				}

				util.AgnosticConfigUpdater("\n# Added by dagger\n")
			}).Run()
		case "fzf":
			_ = spinner.New().Title("Installing fzf...").Action(func() {
				_, err := exec.Command("brew", "install", "fzf").Output()
				if err != nil {
					color.Red("Error installing fzf\n")
					os.Exit(1)
				}
			}).Run()
		case "bat":
			_ = spinner.New().Title("Installing bat...").Action(func() {
				_, err := exec.Command("brew", "install", "bat").Output()
				if err != nil {
					color.Red("Error installing bat\n")
					os.Exit(1)
				}
			}).Run()
		case "ripgrep":
			_ = spinner.New().Title("Installing Ripgrep...").Action(func() {
				_, err := exec.Command("brew", "install", "ripgrep").Output()
				if err != nil {
					color.Red("Error installing ripgrep\n")
					os.Exit(1)
				}
			}).Run()
		case "thefuck":
			_ = spinner.New().Title("Installing thefuck...").Action(func() {
				_, err := exec.Command("brew", "install", "thefuck").Output()
				if err != nil {
					color.Red("Error installing thefuck\n")
					os.Exit(1)
				}

				util.AgnosticConfigUpdater("\n# Added by dagger\n"+ THE_FUCK_SETTING)
			}).Run()
		case "lazygit":
			_ = spinner.New().Title("Installing lazygit..").Action(func() {
				_, err := exec.Command("brew", "install", "lazygit").Output()
				if err != nil {
					color.Red("Error installing lazygit\n")
					os.Exit(1)
				}
			}).Run()
		case "nvm":
			_ = spinner.New().Title("Installing nvm..").Action(func() {
				_, err := exec.Command("brew", "install", "nvm").Output()
				if err != nil {
					color.Red("Error installing nvm\n")
					log.Fatal(err)
					os.Exit(1)
				}

				// create the .nvm folder for the nvm requirements
				exec.Command("mkdir", "~/.nvm")
				util.AgnosticConfigUpdater("\n# Added by dagger\n" + NVM_SH_SETTING)
			}).Run()
		case "skip":
			continue
		}
	}
}

const NVM_SH_SETTING = `export NVM_DIR="$HOME/.nvm"
    [ -s "$HOMEBREW_PREFIX/opt/nvm/nvm.sh" ] && \. "$HOMEBREW_PREFIX/opt/nvm/nvm.sh" # This loads nvm
    [ -s "$HOMEBREW_PREFIX/opt/nvm/etc/bash_completion.d/nvm" ] && \. "$HOMEBREW_PREFIX/opt/nvm/etc/bash_completion.d/nvm" # This loads nvm bash_completion` 

const THE_FUCK_SETTING = "eval $(thefuck --alias)"

const EZA_SETTING = "alias ls='eza --color=always --long --git --no-filesize --no-time --no-user --no-permissions --tree --level=2'"