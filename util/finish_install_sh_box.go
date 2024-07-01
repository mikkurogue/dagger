package util

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	xstrings "github.com/charmbracelet/x/exp/strings"
)

func FinishInstallShBoxMultipleItems(items []string, category string) {
	var sb strings.Builder
	keyword := func(s string) string {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Render(s)
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
}

func FinishInstallShBox(item string, category string) {
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
}
