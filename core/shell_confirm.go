package core

import (
	"github.com/charmbracelet/huh"
)

func ShellConfirm(zsh bool) *huh.Group {
	return huh.NewGroup(
		huh.NewConfirm().
			Title("Is your default shell zsh?").
			Description("This is required for Oh my zsh to function.\n").
			Value(&zsh),
	)
}
