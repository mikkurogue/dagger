package main

import (
	"dagger/core"
	core_windows "dagger/core/windows"
	"dagger/installer"
	installer_windows "dagger/installer/windows"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	xstrings "github.com/charmbracelet/x/exp/strings"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/charmbracelet/lipgloss"
	"github.com/fatih/color"
)

var (
	cli_tools        []string
	aliases          []string
	code_editor      string
	zsh              bool
	zed_installed    bool
	vscode_installed bool
	current_os       string

	curr_step int = 0
)

func CheckOperatingSystem() {
	if runtime.GOOS == "windows" {
		current_os = "windows"
	}

	if runtime.GOOS == "darwin" {
		current_os = "darwin"
	}

	if runtime.GOOS == "linux" {
		current_os = "linux"
	}
}

func main() {

	CheckOperatingSystem()

	if current_os == "windows" {
		WindowsForm()

	} else {
		UnixForm()
	}
}

func WindowsForm() {
	form := huh.NewForm(
		core_windows.Tools(&cli_tools),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	install := func() {
		installer_windows.Tools(cli_tools)
	}

	_ = spinner.New().Title("").Action(install).Run()

	var sb strings.Builder
	keyword := func(s string) string {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Render(s)
	}
	if cli_tools[0] == "skip" {
		fmt.Println(lipgloss.NewStyle().
			Width(40).
			BorderStyle(lipgloss.RoundedBorder()).
			Padding(1, 2).
			Foreground(lipgloss.Color("209")).Render("CLI Tools skipped."))
	} else {
		fmt.Fprintf(&sb,
			"Following tools were installed \n%s\n",
			keyword(xstrings.SpokenLanguageJoin(cli_tools, xstrings.EN)),
		)
		fmt.Println(
			lipgloss.NewStyle().
				Width(40).
				BorderStyle(lipgloss.RoundedBorder()).
				Padding(1, 2).
				Render(sb.String()),
		)
	}
}

func UnixForm() {
	form := huh.NewForm(
		core.Tools(&cli_tools),
		// Install handy dandy aliases
		core.Aliases(&aliases),
		// Install text editor
		core.Editors(&code_editor),
		// Final info about cli installs
		core.ShellConfirm(&zsh),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if !zsh {
		color.Red("Oh my zsh requires zsh as the default shell.\n")
		os.Exit(1)
	}

	install := func() {
		installer.Tools(cli_tools, current_os, curr_step)
		installer.Aliases(aliases, current_os, curr_step)
		installer.CodeEditor(code_editor, current_os, curr_step)
	}

	_ = spinner.New().Title("").Action(install).Run()

	var sb strings.Builder
	keyword := func(s string) string {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Render(s)
	}
	if cli_tools[0] == "skip" {
		fmt.Println(lipgloss.NewStyle().
			Width(40).
			BorderStyle(lipgloss.RoundedBorder()).
			Padding(1, 2).
			Foreground(lipgloss.Color("209")).Render("CLI Tools skipped."))
	} else {
		fmt.Fprintf(&sb,
			"Following tools were installed \n%s\n",
			keyword(xstrings.SpokenLanguageJoin(cli_tools, xstrings.EN)),
		)
		fmt.Println(
			lipgloss.NewStyle().
				Width(40).
				BorderStyle(lipgloss.RoundedBorder()).
				Padding(1, 2).
				Render(sb.String()),
		)
	}

	var aliases_sb strings.Builder
	if aliases[0] == "skip" {
		fmt.Println(lipgloss.NewStyle().
			Width(40).
			BorderStyle(lipgloss.RoundedBorder()).
			Padding(1, 2).
			Foreground(lipgloss.Color("209")).Render("Aliases skipped."))
	} else {
		fmt.Fprintf(&aliases_sb,
			"Following aliases have been set \n%s\n",
			keyword(xstrings.SpokenLanguageJoin(aliases, xstrings.EN)),
		)
		fmt.Println(lipgloss.NewStyle().
			Width(40).
			BorderStyle(lipgloss.RoundedBorder()).
			Padding(1, 2).
			Render(aliases_sb.String()))
	}

	if code_editor == "skip" {
		fmt.Println(lipgloss.NewStyle().
			Width(40).
			BorderStyle(lipgloss.RoundedBorder()).
			Padding(1, 2).
			Foreground(lipgloss.Color("209")).
			Render("Code editor install skipped."))
	} else {
		fmt.Println(lipgloss.NewStyle().
			Width(40).
			BorderStyle(lipgloss.RoundedBorder()).
			Padding(1, 2).
			Foreground(lipgloss.Color("211")).
			Render("Code editor installed " + code_editor))
	}
}