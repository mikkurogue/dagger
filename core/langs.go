package core

import (
	"errors"

	"github.com/charmbracelet/huh"
)

func Langs(langs *[]string) *huh.Group {
	return huh.NewGroup(huh.NewMultiSelect[string]().
		Title("Languages").
		Description("Choose language(s) to install").
		Options(
			huh.NewOption("Go", "go"),
			huh.NewOption("nvm - node version manager", "nvm"),
			huh.NewOption("Skip step", "skip"),
		).
		Validate(func(s []string) error {
			if len(s) == 0 {
				return errors.New("please select at least one tool to install")
			}
			return nil
		}).
		Value(langs))
}
