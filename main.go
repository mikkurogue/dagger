package main

import (
	"dagger/core"
	core_windows "dagger/core/windows"
	"dagger/installer"
	installer_windows "dagger/installer/windows"
	"dagger/util"
	"fmt"
	"log"
	"os"
	"strings"

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

	curr_step int = 0
)

func main() {

	util.DefineOs(&current_os)

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
		core.Langs(&langs),
		core.Aliases(&aliases),
		core.Editors(&code_editor),
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

	util.FinishInstallShBoxMultipleItems(cli_tools, "Tools")
	util.FinishInstallShBoxMultipleItems(langs, "Langs")
	util.FinishInstallShBoxMultipleItems(aliases, "Alias")
	util.FinishInstallShBox(code_editor, "Code editor")

}