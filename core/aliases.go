package core

import (
	"errors"

	"github.com/charmbracelet/huh"
)

func Aliases(aliases *[]string) *huh.Group {
	return huh.NewGroup(huh.NewMultiSelect[string]().
		Title("Aliases").
		Description("Select the aliases you would like to install").
		Options(
			huh.NewOption("git purge", "git-purge"),
			huh.NewOption("Skip step", "skip"),
		).
		Validate(func(s []string) error {
			if len(s) == 0 {
				return errors.New("please select at least one alias to install")
			}
			return nil
		}).
		Value(aliases))
}
