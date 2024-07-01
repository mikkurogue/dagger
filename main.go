package main

import (
	"dagger/config"
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
	cliTools   []string
	aliases    []string
	langs      []string
	codeEditor string
	currentOs  string

	cfg config.Config
)

func main() {
	// update pointer value for the session so we dont need to keep re-assiging

	util.DefineOs(&currentOs)

	config.OpenConfig(&cfg)
}

// Todo: refactor both Form functions to use interfaces and struct methods
func UnixForm() {
	form := huh.NewForm(
		core.Tools(&cliTools),
		core.Langs(&langs),
		core.Aliases(&aliases),
		core.Editors(&codeEditor),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	install := func() {
		installer.Tools(cliTools, currentOs)
		installer.Langs(langs, currentOs)
		installer.Aliases(aliases, currentOs)
		installer.CodeEditor(codeEditor, currentOs)

		cfg := config.Config{}

		cfg.UpdateConfig(aliases, cliTools, codeEditor, langs)
	}

	_ = spinner.New().Title("").TitleStyle(util.TITLE_STYLE).Action(install).Run()

	errTool := util.FinishInstallShBoxMultipleItems(cliTools, "Tools")
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
	errEditor := util.FinishInstallShBox(codeEditor, "Code editor")
	if errEditor != nil {
		log.Fatal(errEditor)
	}
}

func WindowsForm() {
	form := huh.NewForm(
		core_windows.Tools(&cliTools),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	install := func() {
		installer_windows.Tools(cliTools)
	}

	_ = spinner.New().Title("").Action(install).Run()

	toolErr := util.FinishInstallShBoxMultipleItems(cliTools, "Tools")
	if toolErr != nil {
		log.Fatal(toolErr)
	}
}
