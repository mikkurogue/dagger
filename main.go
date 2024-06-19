package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	xstrings "github.com/charmbracelet/x/exp/strings"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/charmbracelet/lipgloss"
	"github.com/fatih/color"
)

var (
	cli_tools []string

	zsh bool
)

func main() {
	form := huh.NewForm(
		//  TODO figure out how to make sure we are authed in git package
		// huh.NewGroup(
		// 	huh.NewSelect[string]().
		// 		Title("Select BigMile Repository to clone").
		// 		Options(
		// 			huh.NewOption("BigMile3 Monorepo", "bigmile3"),
		// 			huh.NewOption("FE Boilerplate", "frontend-boilerplate"), // this is just to test
		// 		).
		// 		Description("Select one of the repositories above to get started.").
		// 		Value(&repo_name)
		// 	huh.NewConfirm().Title("Finish repo setup?"),
		// ),
		huh.NewGroup(huh.NewMultiSelect[string]().
			Title("CLI tools to install").
			Description("Select the tools you would like to install").
			Options(
				huh.NewOption("Eza - better LS", "eza"),
				huh.NewOption("FZF - fuzzy finder", "fzf"),
				huh.NewOption("Bat - better cat", "bat"),
				huh.NewOption("Ripgrep - better grep", "ripgrep"),
				huh.NewOption("Oh my zsh - ZSH theming", "oh-my-zsh"),
				huh.NewOption("TheFuck - CLi typo fixer", "thefuck"),
			).
			Validate(func(s []string) error {
				if len(s) == 0 {
					return errors.New("Please select at least one tool to install")
				}
				return nil
			}).
			Value(&cli_tools)),
		// Final info about cli installs
		huh.NewGroup(
			huh.NewConfirm().
				Title("Is your default shell zsh?").
				Description("This is required for Oh my zsh to function.").
				Value(&zsh),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if !zsh {
		color.Red("Oh my zsh requires zsh as the default shell.")
		os.Exit(1)
	}

	install := func() {
		for _, tool := range cli_tools {
			switch tool {
			case "eza":
				_ = spinner.New().Title("Installing Eza...").Action(func() {
					_, err := exec.Command("brew", "install", "eza").Output()
					if err != nil {
						color.Red("Error installing eza")
						os.Exit(1)
					}

					_, setupErr := exec.Command("/bin/zsh", "-c", "alias ls='eza --color=always --long --git --no-filesize --no-time --no-user --no-permission --tree --level=2'").Output()
					if setupErr != nil {
						color.Red("Can not set ls alias")
					}
				}).Run()
			case "fzf":
				_ = spinner.New().Title("Installing fzf...").Action(func() {
					_, err := exec.Command("brew", "install", "fzf").Output()
					if err != nil {
						color.Red("Error installing fzf")
						os.Exit(1)
					}
				}).Run()
			case "bat":
				_ = spinner.New().Title("Installing bat...").Action(func() {
					_, err := exec.Command("brew", "install", "bat").Output()
					if err != nil {
						color.Red("Error installing bat")
						os.Exit(1)
					}
				}).Run()
			case "ripgrep":
				_ = spinner.New().Title("Installing Ripgrep...").Action(func() {
					_, err := exec.Command("brew", "install", "ripgrep").Output()
					if err != nil {
						color.Red("Error installing ripgrep")
						os.Exit(1)
					}
				}).Run()
			case "oh-my-zsh":
				_ = spinner.New().Title("Installing Oh my zsh...").Action(func() {
					_, err := exec.Command("brew", "install", "oh-my-zsh").Output()
					if err != nil {
						color.Red("Error installing oh my zsh")
						os.Exit(1)
					}
				}).Run()
			case "thefuck":
				_ = spinner.New().Title("Installing thefuck...").Action(func() {
					_, err := exec.Command("brew", "install", "thefuck").Output()
					if err != nil {
						color.Red("Error installing thefuck")
						os.Exit(1)
					}

					_, setupErr := exec.Command("/bin/zsh", "-c", "eval $(thefuck --alias)").Output()
					if setupErr != nil {
						color.Red("Can not set thefuck alias")
					}
				}).Run()
			}
		}
	}

	_ = spinner.New().Action(install).Run()

	var sb strings.Builder
	keyword := func(s string) string {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Render(s)
	}
	fmt.Fprintf(&sb,
		"Following tools were installed \n%s\n",
		keyword(xstrings.EnglishJoin(cli_tools, true)),
	)

	fmt.Println(
		lipgloss.NewStyle().
			Render(sb.String()),
	)
}
