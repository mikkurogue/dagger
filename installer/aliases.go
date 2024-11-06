package installer

import (
	"dagger/util"

	"github.com/charmbracelet/huh/spinner"
	"github.com/fatih/color"
)

func Aliases(aliases []string, currentOs string) {
	for _, alias := range aliases {
		switch alias {
		case "git-purge":
			_ = spinner.New().Title("Setting git-purge alias").Action(func() {
				if currentOs == "windows" {
					color.Red("can not set git-purge alias on this operating system \n")
					return
				}
				util.AgnosticConfigUpdater("\n# Added by dagger\nalias git-purge=\"git fetch -p && git branch --merged | grep -v '*' | grep -v 'master' | xargs git branch -d\"")
			}).TitleStyle(util.TITLE_STYLE).Run()
		case "gmd":
			_ = spinner.New().Title("Setting gmd alias").Action(func() {
				if currentOs == "windows" {
					color.Red("can not set gmd alias on this operating system \n")
				}
				util.AgnosticConfigUpdater("\n# Added by dagger\nalias gmd=\"git merge develop\"")
			}).TitleStyle(util.TITLE_STYLE).Run()
		case "gp":
			_ = spinner.New().Title("Setting gp alias").Action(func() {
				if currentOs == "windows" {
					color.Red("can not set gp alias on this operating system \n")
				}
				util.AgnosticConfigUpdater("\n# Added by dagger\nalias gp=\"git push\"")
			}).TitleStyle(util.TITLE_STYLE).Run()
		case "..":
			_ = spinner.New().Title("Setting .. alias").Action(func() {
				if currentOs == "windows" {
					color.Red("can not set .. alias on this operating system \n")
				}
				util.AgnosticConfigUpdater("\n# Added by dagger\nalias ..=\"cd ..\"")
			}).TitleStyle(util.TITLE_STYLE).Run()
		case "skip":
			continue
		}
	}
}
