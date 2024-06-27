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
	cliTools        []string
	aliases          []string
	langs            []string
	codeEditor      string
	currentOs       string

	cfg config.Config
)

func main() {

	util.DefineOs(&currentOs)

	config.OpenConfig(&cfg)

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

	util.FinishInstallShBoxMultipleItems(cliTools, "Tools")
	util.FinishInstallShBoxMultipleItems(langs, "Langs")
	util.FinishInstallShBoxMultipleItems(aliases, "Alias")
	util.FinishInstallShBox(codeEditor, "Code editor")

}
