package core

import (
	"github.com/charmbracelet/huh"
)

// validation shouldnt be necessary as these aliases arent required.

func Aliases(aliases *[]string) *huh.Group {
	return huh.NewGroup(huh.NewMultiSelect[string]().
		Title("Aliases").
		Description("Select the aliases you would like to install").
		Options(
			huh.NewOption("git purge", "git-purge"),
			huh.NewOption("Skip step", "skip"),
		).
		Value(aliases))
}
