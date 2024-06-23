package main

import (
	"dagger/config"
	"dagger/core"
	"dagger/installer"
	"dagger/util"
	"log"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
)

var (
	cli_tools        []string
	aliases          []string
	langs            []string
	code_editor      string
	zed_installed    bool
	vscode_installed bool
	current_os       string

	cfg config.Config
)

func main() {

	util.DefineOs(&current_os)

	config.OpenConfig(&cfg)

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
