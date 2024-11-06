package core

import (
	"errors"

	"github.com/charmbracelet/huh"
)

func Editors(codeEditor *string) *huh.Group {
	return huh.NewGroup(
		huh.NewSelect[string]().
			Title("Code editor").
			Description("Select the code editor you want to install").
			Options(
				huh.NewOption("Visual Studio Code", "visual-studio-code"),
				huh.NewOption("Zed", "zed"),
				huh.NewOption("NeoVim - expert mode", "neovim"),
				huh.NewOption("Helix", "helix"),
				huh.NewOption("Skip step", "skip"),
			).
			Validate(func(s string) error {
				if s == "" {
					return errors.New("please select at least one code editor to install")
				}
				return nil
			}).
			Value(codeEditor))
}
