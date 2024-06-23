package installer

import (
	"dagger/util"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/charmbracelet/huh/spinner"
	"github.com/charmbracelet/lipgloss"
	"github.com/fatih/color"
)

func Langs(langs []string, current_os string) {
	dir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(lipgloss.NewStyle().
			Background(lipgloss.Color("#ff0000")).
			Foreground(lipgloss.Color("#d3d3d3")).
			Italic(true).
			Padding(0, 1).
			Render("Can not access user home dir"))
	}

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
				nvm_dir := filepath.Join(dir, ".nvm")

				if err := os.MkdirAll(nvm_dir, 0755); err != nil {
					fmt.Println(lipgloss.NewStyle().
						Background(lipgloss.Color("#ff0000")).
						Foreground(lipgloss.Color("#d3d3d3")).
						Italic(true).
						Padding(0, 1).
						Render("Error creating ~/.dagger directory:" + err.Error()))
					return
				}
				util.AgnosticConfigUpdater("\n# Added by dagger\n" + NVM_SH_SETTING)
			}).TitleStyle(util.TITLE_STYLE).Run()
		case "skip":
			continue
		}
	}
}
