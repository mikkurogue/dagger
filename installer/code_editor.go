package installer

import (
	"dagger/util"
	"os"
	"os/exec"

	"github.com/charmbracelet/huh/spinner"
	"github.com/fatih/color"
)

func CodeEditor(codeEditor string, currentOs string) {
	if codeEditor != "skip" {
		if currentOs == "linux" || codeEditor == "zed" {
			color.Red("can not install zed on linux yet...\n")
			codeEditor = "skip"
			return
		} else {
			_ = spinner.New().Title("Installing text editor...").Action(func() {
				_, err := exec.Command("brew", "install", "--cask", codeEditor).Output()
				if err != nil {
					color.Red("Error installing " + codeEditor)
					os.Exit(1)
				}
			}).TitleStyle(util.TITLE_STYLE).Run()
		}
	}
}
