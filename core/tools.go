package core

import (
	"errors"

	"github.com/charmbracelet/huh"
)

func Tools(cli_tools *[]string) *huh.Group {
	return huh.NewGroup(huh.NewMultiSelect[string]().
		Title("CLI tools to install").
		Description("Select the tools you would like to install").
		Options(
			huh.NewOption("Eza - better LS", "eza"),
			huh.NewOption("FZF - fuzzy finder", "fzf"),
			huh.NewOption("Bat - better cat", "bat"),
			huh.NewOption("Ripgrep - better grep", "ripgrep"),
			huh.NewOption("Oh my zsh - ZSH theming", "oh-my-zsh"),
			huh.NewOption("TheFuck - CLI typo fixer", "thefuck"),
			huh.NewOption("lazygit - terminal git manage", "lazygit"),
			huh.NewOption("Skip step", "skip"),
		).
		Validate(func(s []string) error {
			if len(s) == 0 {
				return errors.New("please select at least one tool to install")
			}
			return nil
		}).
		Value(cli_tools))
}
