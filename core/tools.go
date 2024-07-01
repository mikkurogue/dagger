package core

import (
	"errors"

	"github.com/charmbracelet/huh"
)

func Tools(cli_tools *[]string) *huh.Group {
	return huh.NewGroup(huh.NewMultiSelect[string]().
		Title("cli tools to install").
		Description("select the tools you would like to install").
		Options(
			huh.NewOption("Eza - better ls", "eza"),
			huh.NewOption("Zoxide - better cd", "zoxide"),
			huh.NewOption("FZF - fuzzy finder", "fzf"),
			huh.NewOption("Bat - better cat", "bat"),
			huh.NewOption("Ripgrep - better grep", "ripgrep"),
			huh.NewOption("TheFuck - cli typo fixer", "thefuck"),
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
