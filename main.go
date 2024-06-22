package main

import (
	"dagger/core"
	"dagger/installer"
	"dagger/util"
	"fmt"
	"log"
	"os"
	"strings"

	xstrings "github.com/charmbracelet/x/exp/strings"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/charmbracelet/lipgloss"
)

var (
	cli_tools        []string
	aliases          []string
	langs            []string
	code_editor      string
	zed_installed    bool
	vscode_installed bool
	current_os       string
)

func main() {

	util.DefineOs(&current_os)

	form := huh.NewForm(
		core.Tools(&cli_tools),
		core.Langs(&langs),
		// Install handy dandy aliases
		core.Aliases(&aliases),
		// Install text editor
		core.Editors(&code_editor),
		// Final info about cli installs
		// Dont look to confirm the shell anymore - we dont support omz at the moment
		// so bash actually should just work if we change hardcoded zshrc values
		// core.ShellConfirm(&zsh),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	install := func() {
		installer.Tools(cli_tools, current_os)
		installer.Langs(langs, current_os)
		installer.Aliases(aliases, current_os)
		installer.CodeEditor(code_editor, current_os)
	}

	_ = spinner.New().Title("").TitleStyle(util.TITLE_STYLE).Action(install).Run()

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
