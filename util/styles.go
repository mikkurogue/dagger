package util

import "github.com/charmbracelet/lipgloss"

var TITLE_STYLE = lipgloss.NewStyle().
	MarginLeft(1).
	MarginRight(5).
	Padding(0, 1).
	Italic(true).
	Foreground(lipgloss.Color("#FFF7DB"))
