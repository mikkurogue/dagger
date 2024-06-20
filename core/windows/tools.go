package core_windows

import (
	"errors"

	"github.com/charmbracelet/huh"
)

func Tools(cli_tools *[]string) *huh.Group {
	return huh.NewGroup(huh.NewMultiSelect[string]().
		Title("CLI Tools to install (Windows)").
		Description("Select the tools you would like to install").
		Options(
			huh.NewOption("lazygit - terminal git manage", "lazygit"),
		).
		Validate(func (s []string) error {
			if len(s) == 0 {
				return errors.New("please select at least one tool to install")
			}
			return nil
		}).Value(cli_tools))
}