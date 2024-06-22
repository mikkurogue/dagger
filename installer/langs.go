package installer

import (
	"dagger/util"
	"log"
	"os"
	"os/exec"

	"github.com/charmbracelet/huh/spinner"
	"github.com/fatih/color"
)

func Langs(langs []string, current_os string) {
	for _, lang := range langs {
		switch lang {
		case "go":
			_ = spinner.New().Title("Installing Go...").Action(func() {
				_, err := exec.Command("brew", "install", "go").Output()
				if err != nil {
					color.Red("Error installing Go")
					os.Exit(1)
				}
			}).TitleStyle(util.TITLE_STYLE).Run()
		case "nvm":
			_ = spinner.New().Title("Installing nvm..").Action(func() {
				_, err := exec.Command("brew", "install", "nvm").Output()
				if err != nil {
					color.Red("Error installing nvm\n")
					log.Fatal(err)
					os.Exit(1)
				}

				// create the .nvm folder for the nvm requirements
				exec.Command("mkdir", "~/.nvm")
				util.AgnosticConfigUpdater("\n# Added by dagger\n" + NVM_SH_SETTING)
			}).Run()
		case "skip":
			continue
		}
	}
}
