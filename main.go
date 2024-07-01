package main

import (
	"dagger/core"
	core_windows "dagger/core/windows"
	"dagger/installer"
	installer_windows "dagger/installer/windows"
	"dagger/util"
	"log"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
)

var (
	cli_tools   []string
	aliases     []string
	langs       []string
	code_editor string
	current_os  string
)

func main() {

	// update pointer value for the session so we dont need to keep re-assiging
	util.DefineOs(&current_os)

	if current_os == "windows" {
		WindowsForm()
	} else {
		UnixForm()
	}
}

// Todo: refactor both Form functions to use interfaces and struct methods
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

	errTool := util.FinishInstallShBoxMultipleItems(cli_tools, "Tools")
	if errTool != nil {
		log.Fatal(errTool)
	}

	errLang := util.FinishInstallShBoxMultipleItems(langs, "Langs")
	if errLang != nil {
		log.Fatal(errLang)
	}
	errAlias := util.FinishInstallShBoxMultipleItems(aliases, "Alias")
	if errAlias != nil {
		log.Fatal(errAlias)
	}
	errEditor := util.FinishInstallShBox(code_editor, "Code editor")
	if errEditor != nil {
		log.Fatal(errEditor)
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

	toolErr := util.FinishInstallShBoxMultipleItems(cli_tools, "Tools")
	if toolErr != nil {
		log.Fatal(toolErr)
	}
}
