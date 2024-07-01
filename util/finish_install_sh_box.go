package util

import (
	"errors"
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	xstrings "github.com/charmbracelet/x/exp/strings"
)

// For now I couldnt think of a better way to do this so if/else case is used.
// return error when it errors, and nil if its succesful?
func FinishInstallShBoxMultipleItems(items []string, category string) error {
	var sb strings.Builder
	keyword := func(s string) string {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Render(s)
	}

	if len(items) == 0 {
		return errors.New("no selection was made for " + category + " install")
	}

	if items[0] == "skip" {
		fmt.Println(lipgloss.NewStyle().
			Width(40).
			BorderStyle(lipgloss.RoundedBorder()).
			Padding(1, 2).
			Foreground(lipgloss.Color("209")).Render(category + " skipped."))
	} else {
		fmt.Fprintf(&sb,
			"Following %s have been set \n%s\n",
			category,
			keyword(xstrings.SpokenLanguageJoin(items, xstrings.EN)),
		)
		fmt.Println(lipgloss.NewStyle().
			Width(40).
			BorderStyle(lipgloss.RoundedBorder()).
			Padding(1, 2).
			Render(sb.String()))
	}

	return nil
}

// For now I couldnt think of a better way to do this so if/else case is used.
// return error when it errors, and nil if its succesful?
func FinishInstallShBox(item string, category string) error {

	if item == "" {
		return errors.New("no selection was made for " + category + " install")
	}

	if item == "skip" {
		fmt.Println(lipgloss.NewStyle().
			Width(40).
			BorderStyle(lipgloss.RoundedBorder()).
			Padding(1, 2).
			Foreground(lipgloss.Color("209")).
			Render(category + " install skipped."))
	} else {
		fmt.Println(lipgloss.NewStyle().
			Width(40).
			BorderStyle(lipgloss.RoundedBorder()).
			Padding(1, 2).
			Foreground(lipgloss.Color("211")).
			Render(category + " installed " + item))
	}

	return nil
}
